import calendar
import datetime

ORDINALS = { '1st': 0, '2nd': 1, '3rd': 2, '4th': 3, 'last': -1, 'teenth': 'teenth' }

WEEKDAYS = {
  'Monday': calendar.MONDAY,
  'Tuesday': calendar.TUESDAY,
  'Wednesday': calendar.WEDNESDAY,
  'Thursday': calendar.THURSDAY,
  'Friday': calendar.FRIDAY,
  'Saturday': calendar.SATURDAY,
  'Sunday': calendar.SUNDAY
}

def meetup_day(year, month, weekday, ordinal):
  c = calendar.monthcalendar(year, month)
  weekday = WEEKDAYS[weekday]
  ordinal = ORDINALS[ordinal]

  days = [week[weekday] for week in c if week[weekday] != 0]

  if ordinal == 'teenth':
    day = next(day for day in days if day > 12)
  else:
    day = days[ordinal]

  return datetime.date(year, month, day)
