bucket_name = "photo-backup-travis-linkey"
acl = "private"
routes = [
  { 
    path = "GET /",
    authorization = "NONE"
  },
  { 
    path = "GET /buckets/all",
    authorization = "NONE"
  },
  { 
    path = "GET /buckets/objects/all",
    authorization = "NONE"
  },
  { 
    path = "POST /buckets/upload",
    authorization = "NONE"
  },
]
