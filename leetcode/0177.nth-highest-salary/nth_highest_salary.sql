Create table If Not Exists Employee
(
    Id     int,
    Salary int
);
Truncate table Employee;
insert into Employee (Id, Salary)
values ('1', '100');
insert into Employee (Id, Salary)
values ('2', '200');
insert into Employee (Id, Salary)
values ('3', '300');

# Write your MySQL query statement below
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    RETURN (
        SELECT IFNULL((SELECT Salary
                       FROM (SELECT Salary, DENSE_RANK() OVER ( ORDER BY Salary DESC ) AS Id
                             FROM Employee
                             GROUP BY Salary) AS Sort
                       WHERE Sort.Id = N), NULL) AS getNthHighestSalary);
END
