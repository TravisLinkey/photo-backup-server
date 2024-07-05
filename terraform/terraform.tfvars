bucket_name = "photo-backup-travis-linkey"
thumbnail_server_arn = "arn:aws:lambda:us-west-2:731084384219:function:photo-backup-thumbnail-server"
thumbnail_execution_role_arn = "arn:aws:iam::731084384219:role/thumbnail_execution_role"

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
    path = "GET /buckets/presigned-url",
    authorization = "NONE"
  },
  { 
    path = "GET /buckets/thumbnails",
    authorization = "NONE"
  },
]
