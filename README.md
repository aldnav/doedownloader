DOE Downloader
---

A simple downloader for retail prices publicly available at [Department of Energy](https://www.doe.gov.ph/)  

```console
ENV_FILE=~/path-to-env/.env ./bin/downloader
Downloaded a file petro_ncr_2022-may-12.pdf with size 232249
Downloaded a file petro_sluz_2022-may-10_batangas-rizal-quezon.pdf with size 134209
Downloaded a file petro_sluz_2022-may-10_bicol-region.pdf with size 142377
Downloaded a file petro_sluz_2022-may-10_cavite.pdf with size 154803
Downloaded a file petro_sluz_2022-may-10_laguna.pdf with size 123602
Downloaded a file petro_sluz_2022-may-10_mimaropa.pdf with size 161979
Downloaded a file petro_nluz_2022-may-13.pdf with size 335998
Downloaded a file petro-vis_2022-may-10.pdf with size 1198965
Downloaded a file petro_min_2022-may-10.pdf with size 379392
```

Example env file

```env
REPORTS_DIRECTORY=/Users/arthurmorgan/petro/reports
COOKIE_PATH=/Users/arthurmorgan/petro/cookies.json
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


