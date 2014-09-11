slf2tcx
=======

Converts a log file from Sigma Data Center v3 to Garmin TCX file.

	$ ./bin/slf2tcx < 2014_09_04__07_21.slf > workout.tcx

Tested only with log files from Sigma ROX 6.0, no GPS data conversions
are performed.

tcx+gpx
=======

Combines data from a TCX workout and corresponding GPX track, recorded
by some other device. For each track point in TCX at time `t` it
samples GPX track at time `t` and stores position and elevation of the
track in TCX.

	$ cat workout.tcx | ./bin/tcx+gpx 'Track_2014-09-04 072023.gpx' > for_upload_to_strava.tcx

The conversion could be inaccurate:

* Simple linear interpolation is used for sampling track
  points. Probably, it's fine as track points are close to each other.
* Clock in a cycling computer without GPS receiver is likely off from
  the clock in a GPS device. Provided that the difference is
  significant, GPX track will be sampled at wrong points. A possible
  solution would be using cross correlation on elevation data from TCX
  and GPX, finding a proper lag and compensating for it.
