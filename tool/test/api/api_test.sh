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
curl localhost:1323/user/1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac
curl localhost:1323/user/1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac/preferences

echo "------------------------"
curl localhost:1323/breakpoints/1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac/techniques
curl localhost:1323/breakpoints/1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac/history

echo "------------------------"
curl localhost:1323/plans/1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac/current
curl localhost:1323/plans/1a8cc69e-c7d0-4f27-8e2e-a6023c27cdac/plan_history