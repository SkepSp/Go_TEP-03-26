
insert into projects (name) values ('Аналитика');

SELECT p.name, STRING_AGG(t.title, ', ') AS task_list
FROM projects AS p
LEFT JOIN tasks AS t ON t.project_id = p.id
GROUP BY p.name;