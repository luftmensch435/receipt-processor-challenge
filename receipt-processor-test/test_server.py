import requests
import json
import os
import subprocess
import time
import signal

BASE_URL = "http://localhost:8080"

## print test results
def print_result(status, fail_loc=None, exp=None, act=None):
    if status == "FAIL":
        print("FAIL: %s" % fail_loc)
        print("  Expect: %s" % exp)
        print("  Actual: %s" % act)
    else:
        print("PASS")

## test each receipt
def test_receipt(filename, expected_results): 
    exp_submit_status, invalid_id, exp_points = expected_results

    with open(filename, "r") as file:
        receipt = json.load(file)

    # submit receipt
    response = requests.post(f"{BASE_URL}/receipts/process", json=receipt)
    if response.status_code != exp_submit_status:
        print_result("FAIL", "SUBMIT_RECEIPT_STATUS_CODE", exp_submit_status, response.status_code)
        return
    if response.status_code != 200:      
        # PASS, early termination
        print_result("PASS")      
        return

    # test continues, get points
    receipt_id = response.json().get("id")  
    if invalid_id is not None:
        receipt_id = invalid_id
        response = requests.get(f"{BASE_URL}/receipts/{receipt_id}/points")
        if response.status_code != 404:
            print_result("FAIL", "GET_POINTS_STATUS_CODE", 404, response.status_code)
            return
        else:
            # PASS, early termination
            print_result("PASS")      
            return
 
    response = requests.get(f"{BASE_URL}/receipts/{receipt_id}/points")
    if response.status_code != 200:
        print_result("FAIL", "GET_POINTS_STATUS_CODE", 200, response.status_code)
        return

    # check points
    act_points = response.json().get("points")
    if exp_points != act_points:
        print_result("FAIL", "POINTS_VALUE", exp_points, act_points)
        return
    
    # passed
    print_result("PASS")


## Start the Go server before running this test.
if __name__ == "__main__":
    ## expected_results = (exp_submit_status, invalid_id, exp_points)
    ## test_case = (name, filename, expected_results)
    test_cases = [
        ("Bad Input", "./test-receipts/bad-input.json", (400, None, -1)),
        ("Bad UUID", "./test-receipts/readme-receipt-1.json", (200, "", -1)),
        ("Readme case #1", "./test-receipts/readme-receipt-1.json", (200, None, 28)),
        ("Readme case #2", "./test-receipts/readme-receipt-2.json", (200, None, 109)),
    ]

    for i, case in enumerate(test_cases, 1):
        name, filename, expected_results = case
        print("Test Case #%2d [%s]\t" % (i, name), end=" ")
        test_receipt(filename, expected_results)