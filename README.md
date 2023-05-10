### to do service-user-campaign


- Admin req
- [] ~GET Log service
- [] GET Service status

- User Campaign
- [] POST Register
- [] POST Login
- [] POST Check email
- [] POST Check Phone

- Dashboard
- [ ] GET User Profile
- [ ] POST Update_avatar
- [ ] PUT Update User profile

- [ ] POST Logout

# Info

Make database

`migrate create -ext sql -dir database/migrations nama_file_migration`

Run Migrate

```
migrate -database "mysql://root@tcp(127.0.0.1:3306)/service_user_campaign" -path database/migrations up
```
