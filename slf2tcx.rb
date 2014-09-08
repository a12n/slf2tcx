#!/usr/bin/env ruby

require 'happymapper'

#-----------------------------------------------------------------------
# SLF
#-----------------------------------------------------------------------

class GeneralInformation
  include HappyMapper
  tag 'GeneralInformation'
  attribute :log_type, String, :tag => 'logType'
  attribute :serial_number, Integer, :tag => 'serialNumber'
  attribute :unit, String
end

class Log
  include HappyMapper
  attribute :revision, Integer
  element :general_information, GeneralInformation
end

#-----------------------------------------------------------------------
# GPX
#-----------------------------------------------------------------------

#-----------------------------------------------------------------------
# TCX
#-----------------------------------------------------------------------

class Position
  include HappyMapper
  tag 'Position'
  element :latitude, Float, :tag => 'LatitudeDegrees'
  element :longitude, Float, :tag => 'LongitudeDegrees'
end

class HeartRateInBeatsPerMinute
  include HappyMapper
  tag 'HeartRateBpm'
  element :value, Integer, :tag => 'Value'
end

class Trackpoint
  include HappyMapper
  tag 'Trackpoint'
  element :time, Time, :tag => 'Time'
  element :position, Position, :tag => 'Position'
  element :altitude, Float, :tag => 'AltitudeMeters'
  element :distance, Float, :tag => 'DistanceMeters'
  element :heartrate, HeartRateInBeatsPerMinute
  element :cadence, Integer, :tag => 'Cadence'
  element :sensor_state, String, :tag => 'SensorState'
end

class Track
  include HappyMapper
  tag 'Track'
  has_many :trackpoint, Trackpoint, :tag => 'Trackpoint'
end

class ActivityLap
  include HappyMapper
  tag 'Lap'
  attribute :start_time, String, :tag => 'StartTime' # Must be Time
  element :total_time, Float, :tag => 'TotalTimeSeconds'
  element :distance, Float, :tag => 'DistanceMeters'
  element :maximum_speed, Float, :tag => 'MaximumSpeed'
  element :calories, Integer, :tag => 'Calories'
  element :average_heart_rate, HeartRateInBeatsPerMinute, :tag => 'AverageHeartRateBpm'
  element :maximum_heart_rate, HeartRateInBeatsPerMinute, :tag => 'MaximumHeartRateBpm'
  element :intensity, String, :tag => 'Intensity'
  element :cadence, Integer, :tag => 'Cadence'
  element :trigger_method, String, :tag => 'TriggerMethod'
  element :notes, String, :tag => 'Notes'
  has_many :tracks, Track, :tag => 'Track'
end

class Activity
  include HappyMapper
  tag 'Activity'
  attribute :sport, String, :tag => 'Sport'
  element :id, Time, :tag => 'Id'
  has_many :lap, ActivityLap
end

class ActivityList
  include HappyMapper
  tag 'Activities'
  has_many :activity, Activity
end

class TrainingCenterDatabase
  include HappyMapper
  # register_namespace 'tcx', 'http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2'
  # namespace 'tcx'
  tag 'TrainingCenterDatabase'
  element :activities, ActivityList
end

#-----------------------------------------------------------------------
# Main
#-----------------------------------------------------------------------

tcx = TrainingCenterDatabase.parse($stdin.read, :single => true)
print tcx.to_xml
