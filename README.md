# petr-riabinin-finchat-backend-assignment

### How to run the API
`export STRIPE_SK="sk_test_51JS6omAgdvziTvFsG6J8eB21ONPeY5gO5vJPKKk7a4qs0htpWHhDIBTAjMC5JfnWuip9yHEFhAWANPUiAJzZBJ9A00ZG6xJKhR"`\
`go run cmd/main.go`


### How to deploy
[Install and initialize the Cloud SDK.](https://cloud.google.com/sdk/docs/install)\
`gcloud app deploy`


### Deployed web service
[https://powerful-layout-277314.uc.r.appspot.com/](https://powerful-layout-277314.uc.r.appspot.com/)


### Noteable design decisions and improvements.
The service contains two abstractions `routings` and `services` with business logic.\
The must-have improvements could be:
- proper handling errors form Stripe API
- proper requests params validation
- health check
- logging
- unit test

### Manual testing
First of all for each customer we need uniq payment_method_id
```
curl https://api.stripe.com/v1/payment_methods \
  -u sk_test_51JS6omAgdvziTvFsG6J8eB21ONPeY5gO5vJPKKk7a4qs0htpWHhDIBTAjMC5JfnWuip9yHEFhAWANPUiAJzZBJ9A00ZG6xJKhR: \
  -d type=card \
  -d "card[number]"=4242424242424242 \
  -d "card[exp_month]"=8 \
  -d "card[exp_year]"=2022 \
  -d "card[cvc]"=314
```

Other methods and test data combined in the [Postman collection](https://www.getpostman.com/collections/c9c56fd077551149d58f)

### Time capacity
- 3h implementation
- 2h deploy + description




