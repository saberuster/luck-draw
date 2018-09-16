

build:
	@go build luck-draw

first_draw:
	# example usage
	@luck-draw  --src=data/test.csv --to-path=data --num=3 --prize=test_prize
