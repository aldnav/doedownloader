# DOE Downloader

A simple downloader for retail prices publicly available at [Department of Energy](https://www.doe.gov.ph/)  

```console
$ ENV_FILE=~/path-to-env/.env ./bin/downloader
[.] Looking for latest report
[.] Reading from report file: /Users/pro/retailprices/reports/2022-06-21_retail_pump.json
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_ncr_2022-jun-14.pdf with size 235666
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_batangas-rizal-quezon.pdf with size 135598
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_bicol-region.pdf with size 143221
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_cavite.pdf with size 157091
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_laguna.pdf with size 124484
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_mimaropa.pdf with size 164074
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_nluz_2022-jun-17.pdf with size 348019
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_vis_2022-jun-07.pdf with size 798977
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_min_2022-jun-14.pdf with size 243889
[.] Done!
```

## Input

Example env file

```env
REPORTS_DIRECTORY=/Users/pro/retailprices/reports
COOKIE_PATH=/Users/pro/retailprices/tests/cookies.json
```

Example reports file (JSON)

```json
[
  {
    "name": "NCR/Metro Manila",
    "url": "https://www.doe.gov.ph/oil-monitor?q=retail-pump-prices-metro-manila",
    "description": "Prevailing Retail Prices of Petroleum Products in NCR as of May 12, 2022\nDate of Monitoring: May 10-12, 2022",
    "attachments": [
      [
        "petro_ncr_2022-may-12.pdf",
        "https://www.doe.gov.ph/sites/default/files/pdf/price_watch/petro_ncr_2022-may-12.pdf"
      ]
    ]
  }
]
```

## Usage

```console
./bin/downloader
```

Optionally, you can provide a report file as an argument:

```console
./bin/downloader /Users/pro/retailprices/reports/2022-06-21_retail_pump.json
```

## Building

Building natively

```console
go build -o bin/downloader
```

Or using Docker build

```console
DOCKER_BUILDKIT=1 docker build --tag docker-doedownloader-test .
```

## Running

```console
$ ENV_FILE=~/path-to-env/.env ./bin/downloader
[.] Looking for latest report
[.] Reading from report file: /Users/pro/retailprices/reports/2022-06-21_retail_pump.json
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_ncr_2022-jun-14.pdf with size 235666
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_batangas-rizal-quezon.pdf with size 135598
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_bicol-region.pdf with size 143221
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_cavite.pdf with size 157091
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_laguna.pdf with size 124484
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_sluz_2022-jun-14_mimaropa.pdf with size 164074
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_nluz_2022-jun-17.pdf with size 348019
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_vis_2022-jun-07.pdf with size 798977
[.] Downloaded a file /Users/pro/retailprices/reports/2022-06-21/petro_min_2022-jun-14.pdf with size 243889
[.] Done!
```

And in Docker

```console
docker run -v /Users/pro/retailprices/reports:/app/data -v /Users/pro/retailprices/tests/cookies.json:/tmp/cookies.json docker-doedownloader
[.] Looking for latest report
[.] Reading from report file: /app/data/2022-06-21_retail_pump.json
[.] Downloaded a file /app/data/2022-06-21/petro_ncr_2022-jun-14.pdf with size 235666
[.] Downloaded a file /app/data/2022-06-21/petro_sluz_2022-jun-14_batangas-rizal-quezon.pdf with size 135598
[.] Downloaded a file /app/data/2022-06-21/petro_sluz_2022-jun-14_bicol-region.pdf with size 143221
[.] Downloaded a file /app/data/2022-06-21/petro_sluz_2022-jun-14_cavite.pdf with size 157091
[.] Downloaded a file /app/data/2022-06-21/petro_sluz_2022-jun-14_laguna.pdf with size 124484
[.] Downloaded a file /app/data/2022-06-21/petro_sluz_2022-jun-14_mimaropa.pdf with size 164074
[.] Downloaded a file /app/data/2022-06-21/petro_nluz_2022-jun-17.pdf with size 348019
[.] Downloaded a file /app/data/2022-06-21/petro_vis_2022-jun-07.pdf with size 798977
[.] Downloaded a file /app/data/2022-06-21/petro_min_2022-jun-14.pdf with size 243889
[.] Done!
```

## Pitfalls and gotchas

1. The downloaded files are not valid PDF

    Try to check the cookies file. Visit the DOE website and copy every cookie you see.
    Then paste them into the cookies file.
    Then try to download the files again.
