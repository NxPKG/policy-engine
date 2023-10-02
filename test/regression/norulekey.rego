package rules.norulekey

import data.khulnasoft

buckets = khulnasoft.resources("aws_s3_bucket")

bucket_name_contains(bucket, needle) {
	is_string(bucket.bucket)
	contains(bucket.bucket, needle)
}

bucket_name_contains(bucket, needle) {
	is_string(bucket.bucket_prefix)
	contains(bucket.bucket_prefix, needle)
}

make_deny(needle) = ret {
	ret := {info |
		bucket := buckets[_]
		bucket_name_contains(bucket, needle)
		info := {
			"message": sprintf("Bucket names should not contain %s", [needle]),
			"resource": bucket,
		}
	}
}

deny = make_deny("bucket")
