const express = require("express")
const fs = require("fs")
const path = require("path")

const app = express()
const port = process.env.PORT || 3000;
const images = fs.readdirSync('simple_images/hentai/')

app.use("/api/v1/hentai/", express.static(path.join(__dirname, '/simple_images/hentai/')));

app.get("/", (req, res) => {
  res.sendFile(path.join(__dirname, "docs/index.html"));
})
app.get('/api/v1/hentai/random', (req, res) => {
  let chosenFile = images[Math.floor(Math.random() * images.length)]; 

  //res.sendFile(path.join(path.join(__dirname, "simple_images/hentai/"), chosenFile));
  res.json({file : "https://random-good-hanime-api.herokuapp.com/api/v1/hentai/" + chosenFile});
})
app.use(function (req,res,next){
	res.status(404).send('Error 404, page not found');
});

app.listen(port, () => {
  console.log(`Listening at http://localhost:${port}`)
})