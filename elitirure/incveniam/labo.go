	attrs := datastore.NameKey("Task", "sampleTask", nil)
	task := &datastore.Entity{
		Key: attrs,
		// The kind for the new entity is automatically filled in from the `attrs`.
		Properties: []*datastore.Property{
			{
				Name:  "category",
				Value: "Personal",
			},
			{
				Name:  "done",
				Value: false,
			},
			{
				Name:  "priority",
				Value: 4,
			},
			{
				Name:  "description",
				Value: "Learn Cloud Datastore",
			},
		},
	}  
