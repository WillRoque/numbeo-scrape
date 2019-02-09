# Numbeo Scrape

This is a small side project to scrape [numbeo.com](numbeo.com) to calculate how much money one would have left after paying living expenses and rent in cities with relevant amount of data in the website. The data used is always current.

## How To Use

Just [Click Here](https://us-central1-numbeo-scrape.cloudfunctions.net/remaining-money).

That link will do the scraping, generate a spreadsheet with the data and start the download.

The spreadsheet is sorted by **Remaining Money after paying living expenses + rent of a large apartment outside city centre** by default.

##### Obs.: It takes approximately 20 seconds for it to finish running.

## Where Is The Code Running

The code is running in a [Google Cloud Function](https://cloud.google.com/functions).
