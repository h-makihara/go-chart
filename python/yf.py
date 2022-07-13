from yahoo_finance_api2 import share
from yahoo_finance_api2.exceptions import YahooFinanceError

import datetime

my_share = share.Share("8304" + ".T")
symbol_data = None
print(int(datetime.datetime.now().timestamp()))
try:
    symbol_data = my_share.get_historical(share.PERIOD_TYPE_DAY,
                                        100,
                                        share.FREQUENCY_TYPE_HOUR,
                                        1)
except YahooFinanceError as e:
    print(e.message)

print(symbol_data)