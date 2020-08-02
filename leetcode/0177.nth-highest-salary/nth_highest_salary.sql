CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    RETURN (
        SELECT IFNULL((SELECT Salary
                       FROM (SELECT Salary, DENSE_RANK() OVER ( ORDER BY Salary DESC ) AS Id
                             FROM Employee
                             GROUP BY Salary) AS Sort
                       WHERE Sort.Id = N), NULL) AS getNthHighestSalary);
END
