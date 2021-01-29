# create index
CURRENT_PATH=$PWD
curl s -H 'Content-Type:application/json' -X PUT ziyuele.com/test_index_1 -d@ES_STD/add_index.json

# getIndex

curl s ziyuele.com/test_index_1
