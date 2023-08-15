# Поднять лимит в настройках индекса специально для этой задачи
- curl -XPUT -H "Content-Type: application/json" "http://localhost:9200/places/_settings" -d '
{
  "index" : {
    "max_result_window" : 20000
  }
}'



# ex01
http://127.0.0.1:8888/ \
http://127.0.0.1:8888/?page=1 \
http://127.0.0.1:8888/?page=15 \
http://127.0.0.1:8888/?page=1365 \
http://127.0.0.1:8888/?page=9999 \
http://127.0.0.1:8888/?page=-1 \
http://127.0.0.1:8888/?page=foo

# ex02
http://127.0.0.1:8888/api/places/ \
http://127.0.0.1:8888/api/places/?page=1 \
http://127.0.0.1:8888/api/places/?page=1365 \
http://127.0.0.1:8888/api/places/?page=1366 \
http://127.0.0.1:8888/api/places/?page=15 \
http://127.0.0.1:8888/api/places/?page=-1 \
http://127.0.0.1:8888/api/places/?page=foo


# ex04
http://127.0.0.1:8888/api/get_token

# ex03
curl -H "Authorization: Bearer tok" http://127.0.0.1:8888/api/recommend?lat=55.674&lon=37.666

curl -H "Authorization: Bearer {token}" http://127.0.0.1:8888/api/recommend?lat=55.674&lon=37.666 \
curl -H "Authorization: Bearer {token}" http://127.0.0.1:8888/api/recommend \
curl -H "Authorization: Bearer {token}" http://127.0.0.1:8888/api/recommend?lat=foo&lon=foo \
curl -H "Authorization: Bearer {token}" http://127.0.0.1:8888/api/recommend?lat=60&lon=60