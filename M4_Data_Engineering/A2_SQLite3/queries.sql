-- Q1. Write a query to list the names of employees hired after January 1, 2021.

SELECT Name 
FROM Employees 
WHERE HireDate > '2021-01-01';

-- Q2. Write a query to calculate the average salary of employees in each department.

SELECT d.DepartmentName, AVG(e.Salary) AS AverageSalary
FROM Employees e
JOIN Departments d ON e.DepartmentID = d.DepartmentID
GROUP BY d.DepartmentName;


-- Q3. Write a query to find the department name where the total salary is the highest.

SELECT d.DepartmentName, SUM(e.Salary) AS TotalSalary
FROM Employees e
JOIN Departments d ON e.DepartmentID = d.DepartmentID
GROUP BY d.DepartmentName
ORDER BY TotalSalary DESC
LIMIT 1;


-- Q4. Write a query to list all departments that currently have no employees assigned.

SELECT d.DepartmentName 
FROM Departments d
LEFT JOIN Employees e ON d.DepartmentID = e.DepartmentID
WHERE e.EmployeeID IS NULL;


-- Q5. Write a query to fetch all employee details along with their department names.

SELECT e.EmployeeID, e.Name, e.Salary, e.HireDate, d.DepartmentName
FROM Employees e
JOIN Departments d ON e.DepartmentID = d.DepartmentID;

