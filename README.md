Maximize Profit

The package implements an o(n) algorithm that takes an array of stock prices and returns the best profit that could be made from 1 purchase and 1 sale of 1 "Company X" stock yesterday.
The following were the design choices for solving the above problem:
Method 1:
Use two loops. In the outer loop, pick elements one by one and in the inner loop calculate the difference of the picked element with every other element in the array and compare the profit made with the maximum profit calculated so far.
Time Complexity : O(n2)
Space : O(1)

Method 2:
We run through a loop and take the difference with the minimum element found so far. So we need to keep track of 2 things:
1) Maximum difference found so far (maxProfit).
2) Minimum number visited so far (minPrice).
Time Complexity : O(n)
Space : O(1)

Needless to say that we take the Method 2.

install & run
With a correctly configured Go toolchain:
1. go get -u github.com/springwiz/stocks
2. go test ./...