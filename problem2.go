mport "time"


type (

  /*
  GaneshMedical would represent the total number of units of a given medicine available at a point in time. This has the units which are currently available as
  well as the units which are in transit and will be available at a future point in time.
   */
  GaneshMedical struct {
    MedicineName string
    AvailabilityDate time.Time
    ExpiryDate time.Time
    StockUnits int
  }

  /*
  User requests
   */
  MedicinesRequested struct {
    MedicineName string
    UnitRequested int
  }

  /*
  DeliverySlotsForWeek represents all the slots to choose from in a week
   */
  DeliverySlotsForWeek struct {
    Monday OneDay
    Tuesday OneDay
    Wednesday OneDay
    Thrusday OneDay
    Friday OneDay
    Saturday OneDay
    Sunday OneDay
  }

  /*
  OneDay represents the number of slots in a day
   */
  OneDay struct {
    Slot1 Slot
    Slot2 Slot
    Slot3 Slot
    Slot4 Slot
    Slot5 Slot
    Slot6 Slot
  }


  /*
  Slot represents a slot of time. Start time is the start time of the slot. For example for a 4-6 pm slot on Sep 2nd 2019 the startTime would be sometime like
  2019-09-02T14:00:00Z.(RCF3339 Format)
   */
  Slot struct {
    StartTime time.Time
    DisplayTimeSlot string
  }

)



For Each medicine requested, the API will query GaneshMedicals to find all the results where

GaneshMedical.MedicineName = MedicinesRequested.MedicineName &&
MedicinesRequested.UnitRequested <= GaneshMedical.StockUnits

This will get all AvaialabilityDates and ExpireDates for that medicine and units requested. Lets store this is a different struct:
  Results struct {
    AvailabilityDateTime time.Time
    ExpiryDateTime time.Time.
  }

Then for each of these AvailabilityDates, the task is to find the slots.
The API should get all the slots where

Slot.StartTime >= Results.AvailabilityDateTime &&
Slot.StartTime < Results.ExpiryDateTime





