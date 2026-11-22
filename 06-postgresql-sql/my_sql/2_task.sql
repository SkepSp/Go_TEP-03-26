select p.name, count(t.id) as tasks_total from projects as p
left join tasks as t on t.project_id = p.id
group by p.name
