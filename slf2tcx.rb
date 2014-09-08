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

class TrainingCenterDatabase
  include HappyMapper
  register_namespace 'tcx2', 'http://www.garmin.com/xmlschemas/TrainingCenterDatabase/v2'
  namespace 'tcx2'
  tag 'TrainingCenterDatabase'
end

#-----------------------------------------------------------------------
# Main
#-----------------------------------------------------------------------

tcx = TrainingCenterDatabase.new
print tcx.to_xml
