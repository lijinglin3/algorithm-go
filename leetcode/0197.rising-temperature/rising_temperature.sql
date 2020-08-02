SELECT B.Id
FROM Weather A,
     Weather B
WHERE DATEDIFF(B.RecordDate, A.RecordDate) = 1
  AND A.Temperature < B.Temperature
