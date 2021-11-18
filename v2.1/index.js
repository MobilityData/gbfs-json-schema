module.exports = {
  gbfsRequired: true,
  files: options => {
    return [
      { file: 'gbfs_versions', required: false },
      { file: 'system_information', required: true },
      { file: 'vehicle_types', required: false }, // @TODO Conditionally REQUIRED complexe
      { file: 'station_information', required: options.docked },
      { file: 'station_status', required: options.docked },
      { file: 'free_bike_status', required: options.freefloating },
      { file: 'system_hours', required: false },
      { file: 'system_calendar', required: false },
      { file: 'system_regions', required: false },
      { file: 'system_pricing_plans', required: false },
      { file: 'system_alerts', required: false },
      { file: 'geofencing_zones', required: false }
    ]
  }
}
