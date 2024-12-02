#!/bin/bash


# Test User CRUD
echo "Testing User CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe","email":"john@example.com","password":"password","role":"buyer"}' http://localhost:8080/users
curl -X GET http://localhost:8080/users/1
curl -X PUT -H "Content-Type: application/json" -d '{"name":"John Doe Updated","email":"john@example.com","password":"newpassword","role":"buyer"}' http://localhost:8080/users/1
curl -X DELETE http://localhost:8080/users/1

# Test Cart CRUD
echo "Testing Cart CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"buyer_id":1}' http://localhost:8080/carts
curl -X GET http://localhost:8080/carts/1
curl -X PUT -H "Content-Type: application/json" -d '{"buyer_id":1}' http://localhost:8080/carts/1
curl -X DELETE http://localhost:8080/carts/1

# Test Category CRUD
echo "Testing Category CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Electronics","description":"Electronic items"}' http://localhost:8080/categories
curl -X GET http://localhost:8080/categories/1
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Electronics Updated","description":"Updated description"}' http://localhost:8080/categories/1
curl -X DELETE http://localhost:8080/categories/1

# Test Product CRUD
echo "Testing Product CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"name":"Laptop","description":"A powerful laptop","image":"laptop.png","price":1000.00,"stock":10,"seller_id":1,"category_id":1}' http://localhost:8080/products
curl -X GET http://localhost:8080/products/1
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Laptop Updated","description":"Updated description","image":"laptop.png","price":1200.00,"stock":8,"seller_id":1,"category_id":1}' http://localhost:8080/products/1
curl -X DELETE http://localhost:8080/products/1

# Test CartItem CRUD
echo "Testing CartItem CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"cart_id":1,"product_id":1,"quantity":2}' http://localhost:8080/cart_items
curl -X GET http://localhost:8080/cart_items/1
curl -X PUT -H "Content-Type: application/json" -d '{"cart_id":1,"product_id":1,"quantity":3}' http://localhost:8080/cart_items/1
curl -X DELETE http://localhost:8080/cart_items/1

# Test Discussion CRUD
echo "Testing Discussion CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"product_id":1,"user_id":1,"content":"Is this product available?"}' http://localhost:8080/discussions
curl -X GET http://localhost:8080/discussions/1
curl -X PUT -H "Content-Type: application/json" -d '{"product_id":1,"user_id":1,"content":"Updated content"}' http://localhost:8080/discussions/1
curl -X DELETE http://localhost:8080/discussions/1

# Test Migration CRUD
echo "Testing Migration CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"migration":"initial_migration","batch":1}' http://localhost:8080/migrations
curl -X GET http://localhost:8080/migrations/1
curl -X PUT -H "Content-Type: application/json" -d '{"migration":"updated_migration","batch":2}' http://localhost:8080/migrations/1
curl -X DELETE http://localhost:8080/migrations/1

# Test Order CRUD
echo "Testing Order CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"buyer_id":1,"status":"pending","total":100.00}' http://localhost:8080/orders
curl -X GET http://localhost:8080/orders/1
curl -X PUT -H "Content-Type: application/json" -d '{"buyer_id":1,"status":"completed","total":100.00}' http://localhost:8080/orders/1
curl -X DELETE http://localhost:8080/orders/1

# Test OrderItem CRUD
echo "Testing OrderItem CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"order_id":1,"product_id":1,"quantity":1,"price":100.00}' http://localhost:8080/order_items
curl -X GET http://localhost:8080/order_items/1
curl -X PUT -H "Content-Type: application/json" -d '{"order_id":1,"product_id":1,"quantity":2,"price":200.00}' http://localhost:8080/order_items/1
curl -X DELETE http://localhost:8080/order_items/1

# Test Reply CRUD
echo "Testing Reply CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"discussion_id":1,"user_id":1,"content":"Yes, it is available."}' http://localhost:8080/replies
curl -X GET http://localhost:8080/replies/1
curl -X PUT -H "Content-Type: application/json" -d '{"discussion_id":1,"user_id":1,"content":"Updated reply"}' http://localhost:8080/replies/1
curl -X DELETE http://localhost:8080/replies/1

# Test Review CRUD
echo "Testing Review CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"product_id":1,"user_id":1,"rating":5,"comment":"Great product!"}' http://localhost:8080/reviews
curl -X GET http://localhost:8080/reviews/1
curl -X PUT -H "Content-Type: application/json" -d '{"product_id":1,"user_id":1,"rating":4,"comment":"Good product"}' http://localhost:8080/reviews/1
curl -X DELETE http://localhost:8080/reviews/1

# Test Session CRUD
echo "Testing Session CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"id":"session1","user_id":1,"ip_address":"127.0.0.1","user_agent":"Mozilla/5.0","payload":"data","last_activity":1234567890}' http://localhost:8080/sessions
curl -X GET http://localhost:8080/sessions/session1
curl -X PUT -H "Content-Type: application/json" -d '{"id":"session1","user_id":1,"ip_address":"127.0.0.1","user_agent":"Mozilla/5.0","payload":"updated data","last_activity":1234567890}' http://localhost:8080/sessions/session1
curl -X DELETE http://localhost:8080/sessions/session1

# Test Transaction CRUD
echo "Testing Transaction CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"user_id":1,"product_id":1,"quantity":1,"total_price":100.00,"status":"pending"}' http://localhost:8080/transactions
curl -X GET http://localhost:8080/transactions/1
curl -X PUT -H "Content-Type: application/json" -d '{"user_id":1,"product_id":1,"quantity":2,"total_price":200.00,"status":"completed"}' http://localhost:8080/transactions/1
curl -X DELETE http://localhost:8080/transactions/1

# Test Wishlist CRUD
echo "Testing Wishlist CRUD"
curl -X POST -H "Content-Type: application/json" -d '{"user_id":1,"product_id":1}' http://localhost:8080/wishlists
curl -X GET http://localhost:8080/wishlists/1
curl -X PUT -H "Content-Type: application/json" -d '{"user_id":1,"product_id":1}' http://localhost:8080/wishlists/1
curl -X DELETE http://localhost:8080/wishlists/1