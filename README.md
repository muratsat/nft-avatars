# Hash any string to an NFT ape

## Usage

https://github.com/muratsat/nft-avatars/assets/51270744/6e83bb49-c378-4a18-93cc-02a9651a48fe

To get an avatar from string, visit `https://nft-avatars.fly.dev/<your string here>`

For example, below is what you get with `https://nft-avatars.fly.dev/murat`

<p align="center">
  <a href="https://nft-avatars.fly.dev/murat">
	<img src="https://nft-avatars.fly.dev/murat" width="150">
  </a>
</p> 

## How to run:

1. Build docker image:

   `docker build -t nft-avatars:latest`

2. Run docker image:

   `docker run -p 8080:8080 -t nft-avatars:latest`

3. Done, the app will be available at [http://localhost:8080](http://localhost:8080)
