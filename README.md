Openbanking CLI
===============

Interactive and manual command line utility to read data from your bank account.
It uses the [free](https://nordigen.com/en/pricing/) [Nordigen banking API](https://nordigen.com/en/products/account-information/).

[List of supported countries and banks](https://nordigen.com/en/coverage/)

The only you need is to create an account and get a token:
[Nordigen API documentation](https://nordigen.com/en/account_information_documenation/api-documention/overview/)

Interactive Mode
----------------

```shell
❯ openbanking-cli
 ██████  ██████  ███████ ███    ██ ██████   █████  ███    ██ ██   ██ ██ ███    ██  ██████       ██████ ██      ██ 
██    ██ ██   ██ ██      ████   ██ ██   ██ ██   ██ ████   ██ ██  ██  ██ ████   ██ ██           ██      ██      ██ 
██    ██ ██████  █████   ██ ██  ██ ██████  ███████ ██ ██  ██ █████   ██ ██ ██  ██ ██   ███     ██      ██      ██ 
██    ██ ██      ██      ██  ██ ██ ██   ██ ██   ██ ██  ██ ██ ██  ██  ██ ██  ██ ██ ██    ██     ██      ██      ██ 
 ██████  ██      ███████ ██   ████ ██████  ██   ██ ██   ████ ██   ██ ██ ██   ████  ██████       ██████ ███████ ██ 

    

✔ Nordigen Token: ****************************************█
✔ 🇪🇸 Spain (ES)
✔ BBVA (BBVA_BBVAESMM)
Select Max historic days (Default: 730): 730
 SUCCESS  Waiting for authorization...                                                                                                                                                                   
 SUCCESS  Listing accounts...                                                                                                                                                                            
Use the arrow keys to navigate: ↓ ↑ → ← 
? Choose an account: 
  ▸ CUENTA BBVA (*********************)
    CUENTA ONLINE (*********************)
    CUENTA METAS (*********************)

        ------ Account Info ------
        Name: CUENTA BBVA
        Iban: *********************
        Owner: *********************
        Currency: EUR
        Status: enabled
        Created at: 2021-08-01T11:45:58.463499Z
        Last Access: 2021-08-01T11:45:58.606345Z
        Balances:
        
                Amount: 1000 EUR
        
                Amount: 1000 EUR

? What do you want to do?: 
  ▸ Print the list of transactions
    Export transactions to a file
✔ Print the list of transactions
+------------+--------------------------------+--------------------------------+---------+---------+
|    DATE    |             PAYEE              |              MEMO              | OUTFLOW | INFLOW  |
+------------+--------------------------------+--------------------------------+---------+---------+
| 2021-07-30 |  PAGO CON TARJETA              | PAGO CON TARJETA DE SERVICIOS  |   -4.60 |         |
+------------+--------------------------------+--------------------------------+---------+---------+
| 2021-07-29 |  CERVECERIA EL JALEO           | PAGO CON TARJETA EN RESTAURANT |   -2.20 |         |
|            |       ES                       |                                |         |         |
+            +--------------------------------+--------------------------------+---------+---------+
|            | ********************           | BIZUM // OTROS // Coca cola    |   -0.80 |         |
+            +--------------------------------+--------------------------------+---------+---------+
|            |  NICASIO NEVADO JIMENE         | PAGO CON TARJETA DE SERVICIOS  |   -4.70 |         |
|            |       ES                       |                                |         |         |
+------------+--------------------------------+--------------------------------+---------+---------+
| 2021-07-27 |  NICASIO NEVADO JIMENE         | PAGO CON TARJETA DE SERVICIOS  |   -9.40 |         |
|            |       ES                       |                                |         |         |
+            +--------------------------------+--------------------------------+---------+---------+
|            |  CERVECERIA EL JALEO           | PAGO CON TARJETA EN RESTAURANT |   -2.50 |         |
|            |       ES                       |                                |         |         |
+------------+--------------------------------+--------------------------------+---------+---------+
|            | *******************            | TRANSFERENCIAS // TRANSFERENCI |         |    4.00 |
+------------+--------------------------------+--------------------------------+---------+---------+

# 
                        To export an account transactions use: 
        
                        openbanking-cli txns export -a account_id -t {NORDIGEN_TOKEN} -f ofx

```

Command-line Mode
-----------------

```shell
❯ go build main.go -o openbanking-cli

❯ export NORDIGEN_TOKEN = "token"

# LIST COUNTRIES
❯ openbanking-cli countries
+---+----------------+------+
|   |      NAME      | CODE |
+---+----------------+------+
| 🇬🇧 | United Kingdom | GB   |
| 🇩🇪 | Germany        | DE   |
| 🇲🇫 | France         | FR   |
| 🇮🇹 | Italy          | IT   |
| 🇪🇸 | Spain          | ES   |
| 🇵🇹 | Portugal       | PT   |
+---+----------------+------+

# LIST COUNTRY BANKS
❯ openbanking-cli banks -c ES -t $NORDIGEN_TOKEN

 SUCCESS  Listing available ES banks                                                                                                                                                                     
+-----------------------------+--------------------------------+-------------+--------------------+--------------------------------+
|             ID              |              NAME              |     BIC     | MAX  HISTORIC DAYS |           COUNTRIES            |
+-----------------------------+--------------------------------+-------------+--------------------+--------------------------------+
| ARQUIA_CASDESBB             | Arquia Banca                   | CASDESBB    |                730 | [ES]                           |
| BBVA_BBVAESMM               | BBVA                           | BBVAESMM    |                730 | [ES]                           |
| MARCH_BMARES2M              | Banca March                    | BMARES2M    |                730 | [ES]                           |
| MEDIOLANUM_BFIVESBB         | Banco Mediolanum               | BFIVESBB    |                365 | [ES]                           |
| PICHINCHA_PICHESMM          | Banco Pichincha                | PICHESMM    |                730 | [ES]                           |
| SANTANDER_BSCHESMM          | Banco Santander                | BSCHESMM    |                730 | [ES]                           |
| BIGES_IGSEESMM              | Banco de Investimento Global   | IGSEESMM    |                 90 | [ES]                           |
| BANCSABADELL_BSABESBB       | Banco de Sabadell              | BSABESBB    |                730 | [ES]                           |
| BANKIA_CAHMESMM             | Bankia                         | CAHMESMM    |                700 | [ES]                           |
| BANKINTER_BKBKESMM          | Bankinter                      | BKBKESMM    |                540 | [ES]                           |
| BANKOA_BKOAES22             | Bankoa                         | BKOAES22    |                730 | [ES]                           |
| CAIXABANK_CAIXESBB          | CaixaBank                      | CAIXESBB    |                730 | [ES]                           |
| CAJARURAL_BCOEESMM          | Caja Rural                     | BCOEESMM    |                730 | [ES]                           |
| CAJASUR_CSURES2C            | Cajasur                        | CSURES2C    |                730 | [ES]                           |
| COLONYA_CECAESMM056         | Colonya                        | CECAESMM056 |                730 | [ES]                           |
| EVOBANCO_EVOBESMM           | EVO Banco                      | EVOBESMM    |                730 | [ES]                           |
| EUROCAJARURAL_BCOEESMM081   | Eurocaja Rural                 | BCOEESMM081 |                730 | [ES]                           |
| FIARE_ETICES21              | Fiare Banca Etica              | ETICES21    |                730 | [ES]                           |
| CAJAMAR_BCCAESMM            | Grupo Cajamar                  | BCCAESMM    |                365 | [ES]                           |
| ING_INGDESMM                | ING                            | INGDESMM    |                730 | [ES]                           |
| IBERCAJA_CAZRES2Z           | IberCaja                       | CAZRES2Z    |                730 | [ES]                           |
| KUTXABANK_BASKES2B          | Kutxabank                      | BASKES2B    |                730 | [ES]                           |
| LABORALKUTXA_CLPEES2M       | Laboral Kutxa                  | CLPEES2M    |                730 | [ES]                           |
| N26_NTSBDEB1                | N26 Bank                       | NTSBDEB1    |                730 | [NO SE FI DK EE LV LT GB NL CZ |
|                             |                                |             |                    | ES PL BE DE AT BG HR CY FR GR  |
|                             |                                |             |                    | HU IS IE IT LI LU MT PT RO SK  |
|                             |                                |             |                    | SI]                            |
| OPENBANK_OPENESMM           | Openbank                       | OPENESMM    |                730 | [ES]                           |
| PAYPAL_PPLXLULL             | PayPal                         | PPLXLULL    |                 90 | [NO SE FI DK EE LV LT GB NL CZ |
|                             |                                |             |                    | ES PL BE DE AT BG HR CY FR GR  |
|                             |                                |             |                    | HU IS IE IT LI LU MT PT RO SK  |
|                             |                                |             |                    | SI]                            |
| RENTA4_RENBESMM             | Renta 4 Banco                  | RENBESMM    |                730 | [ES]                           |
| REVOLUT_REVOGB21            | Revolut                        | REVOGB21    |                730 | [NO SE FI DK EE LV LT GB NL CZ |
|                             |                                |             |                    | ES PL BE DE AT BG HR CY FR GR  |
|                             |                                |             |                    | HU IS IE IT LI LU MT PT RO SK  |
|                             |                                |             |                    | SI]                            |
| CORPORATESANTANDER_BSCHXXMM | Santander Corporate &          | BSCHXXMM    |                730 | [ES]                           |
|                             | Investment Banking             |             |                    |                                |
| UNICAJA_UCJAES2M            | Unicaja Banco                  | UCJAES2M    |                730 | [ES]                           |
| WIZINK_POPLESMM             | WiZink                         | POPLESMM    |                365 | [ES]                           |
| WISE_TRWIGB22               | Wise                           | TRWIGB22    |                730 | [NO SE FI DK EE LV LT GB NL CZ |
|                             |                                |             |                    | ES PL BE DE AT BG HR CY FR GR  |
|                             |                                |             |                    | HU IS IE IT LI LU MT PT RO SK  |
|                             |                                |             |                    | SI]                            |
+-----------------------------+--------------------------------+-------------+--------------------+--------------------------------+


# AUTHORIZE WITH YOUR BANK CREDENTIALS
❯ openbanking-cli auth -b BBVA_BBVAESMM -t $NORDIGEN_TOKEN

 SUCCESS  Getting Credentials                                                                                                                                                                            

Your authorization id : $CREDENTIAL_ID

# LIST ACCOUNTS
❯ openbanking-cli accounts -b BBVAESMM -r $CREDENTIAL_ID -t $NORDIGEN_TOKEN


 SUCCESS  Listing accounts...                                                                                                                                                                            
+----------------------------------------+--------------------------+---------+----------------------+-------------------------------+
|               ID & NAME                |           IBAN           | STATUS  |        OWNER         |            BALANCE            |
+----------------------------------------+--------------------------+---------+----------------------+-------------------------------+
| CUENTA  BBVA                           | *******************      | enabled | NAME                 | closingBooked = 1000EUR       |
| (account_id1)                          |                          |         |                      |                               |
+                                        +                          +         +                      +-------------------------------+
|                                        |                          |         |                      | interimAvailable = 1000EUR    |
|                                        |                          |         |                      |                               |
+----------------------------------------+--------------------------+         +                      +-------------------------------+
| CUENTA ONLINE                          | *******************      |         |                      | closingBooked = 1000EUR       |
| (id2)                                  |                          |         |                      |                               |
+                                        +                          +         +                      +-------------------------------+
|                                        |                          |         |                      | interimAvailable = 1000EUR    |
|                                        |                          |         |                      |                               |
+----------------------------------------+--------------------------+         +                      +-------------------------------+
| CUENTA METAS                           | *******************      |         |                      | closingBooked = 0.00EUR       |
| (id3)                                  |                          |         |                      |                               |
+                                        +                          +         +                      +-------------------------------+
|                                        |                          |         |                      | interimAvailable = 0.00EUR    |
|                                        |                          |         |                      |                               |
+----------------------------------------+--------------------------+---------+----------------------+-------------------------------+

# LIST RECENT TRANSACTIONS
❯ openbanking-cli txns -a account_id1 -t $NORDIGEN_TOKEN

▄ Listing Transactions for account account_id1
+------------+--------------------------------+--------------------------------+---------+---------+
|    DATE    |             PAYEE              |              MEMO              | OUTFLOW | INFLOW  |
+------------+--------------------------------+--------------------------------+---------+---------+
| 2021-07-30 |  PAGO CON TARJETA              | PAGO CON TARJETA DE SERVICIOS  |   -4.60 |         |
+------------+--------------------------------+--------------------------------+---------+---------+
| 2021-07-29 |  CERVECERIA EL JALEO           | PAGO CON TARJETA EN RESTAURANT |   -2.20 |         |
|            |              ES                |                                |         |         |
+            +--------------------------------+--------------------------------+---------+---------+
|            | NAME                           | BIZUM // OTROS // Coca cola    |   -0.80 |         |
+            +--------------------------------+--------------------------------+---------+---------+
|            |  NICASIO NEVADO JIMENE         | PAGO CON TARJETA DE SERVICIOS  |   -4.70 |         |
|            |              ES                |                                |         |         |
+------------+--------------------------------+--------------------------------+---------+---------+

 SUCCESS  Listing Transactions for account account_id1
 
# EXPORT RECENT TRANSACTIONS
❯ openbanking-cli txns export -a account_id1 -t $NORDIGEN_TOKEN -o ofx > txns.ofx
❯ openbanking-cli txns export -a account_id1 -t $NORDIGEN_TOKEN -o csv > txns.csv
❯ openbanking-cli txns export -a account_id1 -t $NORDIGEN_TOKEN -o json > txns.json
 
```