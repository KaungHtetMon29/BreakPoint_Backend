insert into break_point_techniques ("user_uuid","is_active","created_at","technique")
values ('1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac',
true,CURRENT_TIMESTAMP,
'{"theme": "dark", "notifications": {"email": true, "sms": false}}'::jsonb);


insert into break_point_generate_histories ("user_uuid","created_at")
values ('1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac',CURRENT_TIMESTAMP);