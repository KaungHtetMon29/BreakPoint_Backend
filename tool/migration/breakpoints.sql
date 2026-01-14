insert into break_point_techniques ("user_uuid","is_active","created_at","technique")
values ('3b911a22-9e99-4d05-afec-c3f25bc66f15',
true,CURRENT_TIMESTAMP,
'{"theme": "dark", "notifications": {"email": true, "sms": false}}'::jsonb);


insert into break_point_generate_histories ("user_uuid","created_at")
values ('3b911a22-9e99-4d05-afec-c3f25bc66f15',CURRENT_TIMESTAMP);