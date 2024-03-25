docker build -t emojidb .
docker run --name emojidb -p 5432:15200 --rm emojidb