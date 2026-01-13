curl localhost:1323/plans/current
curl localhost:1323/plans/plan_history
curl localhost:1323/plans/usage

echo "------------------------"

curl -X POST localhost:1323/admin/add_admin
curl -X POST localhost:1323/admin/auth/login
curl -X POST localhost:1323/admin/auth/logout
curl -X GET localhost:1323/admin/users
curl -X GET localhost:1323/admin/users/{id}
curl -X PATCH localhost:1323/admin/users/{id}/status
curl -X DELETE localhost:1323/admin/{adminId}
curl -X GET localhost:1323/admin/{adminId}
curl -X PUT localhost:1323/admin/{adminId}
echo "------------------------"
curl localhost:1323/user/3b911a22-9e99-4d05-afec-c3f25bc66f15
curl localhost:1323/user/3b911a22-9e99-4d05-afec-c3f25bc66f15/preferences
