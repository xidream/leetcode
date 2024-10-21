# 2. Add Two Numbers

<a href="https://github.com/whateverzpy/LeetCode-Markdown"><img src="https://img.shields.io/badge/Markdown-FFA116?logo=leetcode&labelColor=555"/></a>
![](https://img.shields.io/badge/Difficulty-Medium-orange)
![](https://img.shields.io/badge/Topics-Linked_List,_Math,_Recursion-blue)

You are given two **non\-empty**  linked lists representing two non\-negative integers\. The digits are stored in **reverse order** , and each of their nodes contains a single digit\. Add the two numbers and return the sum as a linked list\.

You may assume the two numbers do not contain any leading zero, except the number 0 itself\.

<br>

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)
> **Input:**  l1 = \[2,4,3], l2 = \[5,6,4] <br>
> **Output:**  \[7,0,8] <br>
> **Explanation:**  342 \+ 465 = 807\. <br>

**Example 2:**

> **Input:**  l1 = \[0], l2 = \[0] <br>
> **Output:**  \[0] <br>

**Example 3:**

> **Input:**  l1 = \[9,9,9,9,9,9,9], l2 = \[9,9,9,9] <br>
> **Output:**  \[8,9,9,9,0,0,0,1] <br>

<br>

**Constraints:**

*   The number of nodes in each linked list is in the range `[1, 100]`\.
*   `0 <= Node.val <= 9`
*   It is guaranteed that the list represents a number that does not have leading zeros\.