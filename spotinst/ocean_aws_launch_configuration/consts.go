package ocean_aws_launch_configuration

import "github.com/spotinst/terraform-provider-spotinst/spotinst/commons"

const (
	ImageID                  commons.FieldName = "image_id"
	IAMInstanceProfile       commons.FieldName = "iam_instance_profile"
	KeyName                  commons.FieldName = "key_name"
	UserData                 commons.FieldName = "user_data"
	SecurityGroups           commons.FieldName = "security_groups"
	AssociatePublicIpAddress commons.FieldName = "associate_public_ip_address"
	LoadBalancers            commons.FieldName = "load_balancers"
	Arn                      commons.FieldName = "arn"
	Name                     commons.FieldName = "name"
	Type                     commons.FieldName = "type"
	RootVolumeSize           commons.FieldName = "root_volume_size"
	Monitoring               commons.FieldName = "monitoring"
	EBSOptimized             commons.FieldName = "ebs_optimized"
	UseAsTemplateOnly        commons.FieldName = "use_as_template_only"
)

const (
	InstanceMetadataOptions commons.FieldName = "instance_metadata_options"
	HTTPTokens              commons.FieldName = "http_tokens"
	HTTPPutResponseHopLimit commons.FieldName = "http_put_response_hop_limit"
)
