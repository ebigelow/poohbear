from datetime import *
import trade_pb2

def rfc3339_time(dt):
    if dt.utcoffset() is not None:
        return dt.isoformat()
    else:
        return "%sZ" % dt.isoformat()

def up_to_now(exchange="", pair="", days_ago=0,
              hours_ago=0, minutes_ago=0, seconds_ago=0):

    end = rfc3339_time(datetime.utcnow())

    delta = timedelta(days=days_ago,
                      hours=hours_ago,
                      minutes=minutes_ago,
                      seconds=seconds_ago)

    start = rfc3339_time(datetime.utcnow() - delta)

    time_range = trade_pb2.DateRange(exchange=exchange,
                                     pair=pair,
                                     start=start,
                                     end=end)

    return time_range
