---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "cloudautomator_job Data Source - terraform-provider-cloudautomator"
subcategory: ""
description: |-

---

# cloudautomator_job (Data Source)

## Example Usage

```hcl
data "cloudautomator_job" "example-job" {
  id = 123
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) ジョブID

### Read-Only

- `action_type` (String) アクションのタイプ
- `allow_runtime_action_values` (Boolean) アクションの設定値を実行時に指定するかどうか
- `authorize_security_group_ingress_action_value` (List of Object) 「EC2: セキュリティグループにインバウンドルールを追加」アクションの設定値 (see [below for nested schema](#nestedatt--authorize_security_group_ingress_action_value))
- `aws_account_id` (Number) AWSアカウントID
- `change_instance_type_action_value` (List of Object) 「EC2: インスタンスタイプを変更」アクションの設定値 (see [below for nested schema](#nestedatt--change_instance_type_action_value))
- `change_rds_cluster_instance_class_action_value` (List of Object) 「RDS(Aurora): DBインスタンスクラスを変更」アクションの設定値 (see [below for nested schema](#nestedatt--change_rds_cluster_instance_class_action_value))
- `change_rds_instance_class_action_value` (List of Object) 「RDS: DBインスタンスクラスを変更」アクションの設定値 (see [below for nested schema](#nestedatt--change_rds_instance_class_action_value))
- `completed_post_process_id` (List of Number) ジョブが成功した場合に実行する後処理IDが含まれる配列
- `copy_ebs_snapshot_action_value` (List of Object) 「EC2: EBSスナップショットをリージョン間でコピー」アクションの設定値 (see [below for nested schema](#nestedatt--copy_ebs_snapshot_action_value))
- `copy_image_action_value` (List of Object) 「EC2: AMIをリージョン間でコピー」アクションの設定値 (see [below for nested schema](#nestedatt--copy_image_action_value))
- `copy_rds_snapshot_action_value` (List of Object) 「RDS: DBスナップショットをリージョン間でコピー」アクションの設定値 (see [below for nested schema](#nestedatt--copy_rds_snapshot_action_value))
- `create_ebs_snapshot_action_value` (List of Object) 「EC2: EBSスナップショットを作成」アクションの設定値 (see [below for nested schema](#nestedatt--create_ebs_snapshot_action_value))
- `create_image_action_value` (List of Object) 「EC2: AMIを作成」アクションの設定値 (see [below for nested schema](#nestedatt--create_image_action_value))
- `create_rds_cluster_snapshot_action_value` (List of Object) 「RDS(Aurora): DBクラスタースナップショットを作成」アクションの設定値 (see [below for nested schema](#nestedatt--create_rds_cluster_snapshot_action_value))
- `create_rds_snapshot_action_value` (List of Object) 「RDS: DBスナップショットを作成」アクションの設定値 (see [below for nested schema](#nestedatt--create_rds_snapshot_action_value))
- `create_redshift_snapshot_action_value` (List of Object) 「Redshift: クラスタースナップショットを作成」アクションの設定値 (see [below for nested schema](#nestedatt--create_redshift_snapshot_action_value))
- `cron_rule_value` (List of Object) タイマートリガーの設定値 (see [below for nested schema](#nestedatt--cron_rule_value))
- `delay_action_value` (List of Object) 「Other: 指定時間待機」アクションの設定値 (see [below for nested schema](#nestedatt--delay_action_value))
- `delete_cluster_action_value` (List of Object) 「Redshift: クラスターを削除」アクションの設定値 (see [below for nested schema](#nestedatt--delete_cluster_action_value))
- `delete_rds_cluster_action_value` (List of Object) 「RDS(Aurora): DBクラスターを削除」アクションの設定値 (see [below for nested schema](#nestedatt--delete_rds_cluster_action_value))
- `delete_rds_instance_action_value` (List of Object) 「RDS: DBインスタンスを削除」アクションの設定値 (see [below for nested schema](#nestedatt--delete_rds_instance_action_value))
- `deregister_instances_action_value` (List of Object) 「ELB(CLB): EC2インスタンスを登録解除」アクションの設定値 (see [below for nested schema](#nestedatt--deregister_instances_action_value))
- `deregister_target_instances_action_value` (List of Object) 「ELB(ALB/NLB): ターゲットグループからEC2インスタンスを登録解除」アクションの設定値 (see [below for nested schema](#nestedatt--deregister_target_instances_action_value))
- `effective_date` (String) ジョブの有効期間の開始日
- `expiration_date` (String) ジョブの有効期間の終了日
- `failed_post_process_id` (List of Number) ジョブが失敗した場合に実行する後処理IDが含まれる配列
- `group_id` (Number) グループID
- `name` (String) ジョブ名
- `reboot_rds_instances_action_value` (List of Object) 「RDS: DBインスタンスを再起動」アクションの設定値 (see [below for nested schema](#nestedatt--reboot_rds_instances_action_value))
- `reboot_workspaces_action_value` (List of Object) 「WorkSpaces: WorkSpaceを再起動」アクションの設定値 (see [below for nested schema](#nestedatt--reboot_workspaces_action_value))
- `register_instances_action_value` (List of Object) 「ELB(CLB): EC2インスタンスを登録」アクションの設定値 (see [below for nested schema](#nestedatt--register_instances_action_value))
- `register_target_instances_action_value` (List of Object) 「ELB(ALB/NLB): ターゲットグループにEC2インスタンスを登録」アクションの設定値 (see [below for nested schema](#nestedatt--register_target_instances_action_value))
- `restore_from_cluster_snapshot_action_value` (List of Object) 「Redshift: スナップショットからリストア」アクションの設定値 (see [below for nested schema](#nestedatt--restore_from_cluster_snapshot_action_value))
- `restore_rds_cluster_action_value` (List of Object) 「RDS(Aurora): DBクラスタースナップショットからリストア」アクションの設定値 (see [below for nested schema](#nestedatt--restore_rds_cluster_action_value))
- `restore_rds_instance_action_value` (List of Object) 「RDS: DBスナップショットからリストア」アクションの設定値 (see [below for nested schema](#nestedatt--restore_rds_instance_action_value))
- `revoke_security_group_ingress_action_value` (List of Object) 「EC2: セキュリティグループからインバウンドルールを削除」アクションの設定値 (see [below for nested schema](#nestedatt--revoke_security_group_ingress_action_value))
- `rule_type` (String) トリガーのタイプ
- `schedule_rule_value` (List of Object) スケジュールトリガーの設定値 (see [below for nested schema](#nestedatt--schedule_rule_value))
- `send_command_action_value` (List of Object) 「EC2: インスタンスでコマンドを実行」アクションの設定値 (see [below for nested schema](#nestedatt--send_command_action_value))
- `sqs_v2_rule_value` (List of Object) SQSトリガーの設定値 (see [below for nested schema](#nestedatt--sqs_v2_rule_value))
- `start_instances_action_value` (List of Object) 「EC2: インスタンスを起動」アクションの設定値 (see [below for nested schema](#nestedatt--start_instances_action_value))
- `start_rds_clusters_action_value` (List of Object) 「RDS(Aurora): DBクラスターを起動」アクションの設定値 (see [below for nested schema](#nestedatt--start_rds_clusters_action_value))
- `start_rds_instances_action_value` (List of Object) 「RDS: DBインスタンスを起動」アクションの設定値 (see [below for nested schema](#nestedatt--start_rds_instances_action_value))
- `start_workspaces_action_value` (List of Object) 「WorkSpaces: WorkSpaceを起動」アクションの設定値 (see [below for nested schema](#nestedatt--start_workspaces_action_value))
- `stop_instances_action_value` (List of Object) 「EC2: インスタンスを停止」アクションの設定値 (see [below for nested schema](#nestedatt--stop_instances_action_value))
- `stop_rds_clusters_action_value` (List of Object) 「RDS(Aurora): DBクラスターを停止」アクションの設定値 (see [below for nested schema](#nestedatt--stop_rds_clusters_action_value))
- `stop_rds_instances_action_value` (List of Object) 「RDS: DBインスタンスを停止」アクションの設定値 (see [below for nested schema](#nestedatt--stop_rds_instances_action_value))
- `terminate_workspaces_action_value` (List of Object) 「WorkSpaces: WorkSpaceを削除」アクションの設定値 (see [below for nested schema](#nestedatt--terminate_workspaces_action_value))
- `update_record_set_action_value` (List of Object) 「Route 53: リソースレコードセットを更新」アクションの設定値 (see [below for nested schema](#nestedatt--update_record_set_action_value))
- `windows_update_action_value` (List of Object) 「EC2: インスタンスをWindows Update」アクションの設定値 (see [below for nested schema](#nestedatt--windows_update_action_value))
- `windows_update_v2_action_value` (List of Object) 「EC2: インスタンスをWindows Update (新バージョン)」アクションの設定値 (see [below for nested schema](#nestedatt--windows_update_v2_action_value))

<a id="nestedatt--authorize_security_group_ingress_action_value"></a>
### Nested Schema for `authorize_security_group_ingress_action_value`

Read-Only:

- `cidr_ip` (String)
- `ip_protocol` (String)
- `region` (String)
- `security_group_id` (String)
- `specify_security_group` (String)
- `tag_key` (String)
- `tag_value` (String)
- `to_port` (String)


<a id="nestedatt--change_instance_type_action_value"></a>
### Nested Schema for `change_instance_type_action_value`

Read-Only:

- `instance_id` (String)
- `instance_type` (String)
- `region` (String)
- `specify_instance` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--change_rds_cluster_instance_class_action_value"></a>
### Nested Schema for `change_rds_cluster_instance_class_action_value`

Read-Only:

- `db_instance_class` (String)
- `rds_instance_id` (String)
- `region` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--change_rds_instance_class_action_value"></a>
### Nested Schema for `change_rds_instance_class_action_value`

Read-Only:

- `db_instance_class` (String)
- `rds_instance_id` (String)
- `region` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--copy_ebs_snapshot_action_value"></a>
### Nested Schema for `copy_ebs_snapshot_action_value`

Read-Only:

- `destination_region` (String)
- `snapshot_id` (String)
- `source_region` (String)
- `specify_ebs_snapshot` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--copy_image_action_value"></a>
### Nested Schema for `copy_image_action_value`

Read-Only:

- `destination_region` (String)
- `source_image_id` (String)
- `source_region` (String)
- `specify_image` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--copy_rds_snapshot_action_value"></a>
### Nested Schema for `copy_rds_snapshot_action_value`

Read-Only:

- `destination_region` (String)
- `option_group_name` (String)
- `rds_snapshot_id` (String)
- `source_rds_instance_id` (String)
- `source_region` (String)
- `specify_rds_snapshot` (String)
- `trace_status` (String)


<a id="nestedatt--create_ebs_snapshot_action_value"></a>
### Nested Schema for `create_ebs_snapshot_action_value`

Read-Only:

- `additional_tag_key` (String)
- `additional_tag_value` (String)
- `description` (String)
- `generation` (Number)
- `region` (String)
- `specify_volume` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)
- `volume_id` (String)


<a id="nestedatt--create_image_action_value"></a>
### Nested Schema for `create_image_action_value`

Read-Only:

- `add_same_tag_to_snapshot` (String)
- `additional_tag_key` (String)
- `additional_tag_value` (String)
- `additional_tags` (Set of Object) (see [below for nested schema](#nestedobjatt--create_image_action_value--additional_tags))
- `description` (String)
- `generation` (Number)
- `image_name` (String)
- `instance_id` (String)
- `reboot_instance` (String)
- `recreate_image_if_ami_status_failed` (String)
- `region` (String)
- `specify_image_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)

<a id="nestedobjatt--create_image_action_value--additional_tags"></a>
### Nested Schema for `create_image_action_value.additional_tags`

Read-Only:

- `key` (String)
- `value` (String)



<a id="nestedatt--create_rds_cluster_snapshot_action_value"></a>
### Nested Schema for `create_rds_cluster_snapshot_action_value`

Read-Only:

- `db_cluster_identifier` (String)
- `db_cluster_snapshot_identifier` (String)
- `generation` (Number)
- `region` (String)
- `specify_rds_cluster` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--create_rds_snapshot_action_value"></a>
### Nested Schema for `create_rds_snapshot_action_value`

Read-Only:

- `generation` (Number)
- `rds_instance_id` (String)
- `rds_snapshot_id` (String)
- `region` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--create_redshift_snapshot_action_value"></a>
### Nested Schema for `create_redshift_snapshot_action_value`

Read-Only:

- `cluster_identifier` (String)
- `cluster_snapshot_identifier` (String)
- `generation` (Number)
- `region` (String)
- `specify_cluster` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--cron_rule_value"></a>
### Nested Schema for `cron_rule_value`

Read-Only:

- `dates_to_skip` (List of String)
- `hour` (String)
- `minutes` (String)
- `monthly_day_of_week_schedule` (List of Object) (see [below for nested schema](#nestedobjatt--cron_rule_value--monthly_day_of_week_schedule))
- `monthly_schedule` (String)
- `national_holiday_schedule` (String)
- `one_time_schedule` (String)
- `schedule_type` (String)
- `start_timeout_minutes` (String)
- `time_zone` (String)
- `weekly_schedule` (List of String)

<a id="nestedobjatt--cron_rule_value--monthly_day_of_week_schedule"></a>
### Nested Schema for `cron_rule_value.monthly_day_of_week_schedule`

Read-Only:

- `friday` (List of Number)
- `monday` (List of Number)
- `saturday` (List of Number)
- `sunday` (List of Number)
- `thursday` (List of Number)
- `tuesday` (List of Number)
- `wednesday` (List of Number)



<a id="nestedatt--delay_action_value"></a>
### Nested Schema for `delay_action_value`

Read-Only:

- `delay_minutes` (Number)


<a id="nestedatt--delete_cluster_action_value"></a>
### Nested Schema for `delete_cluster_action_value`

Read-Only:

- `cluster_identifier` (String)
- `final_cluster_snapshot_identifier` (String)
- `region` (String)
- `skip_final_cluster_snapshot` (String)
- `trace_status` (String)


<a id="nestedatt--delete_rds_cluster_action_value"></a>
### Nested Schema for `delete_rds_cluster_action_value`

Read-Only:

- `db_cluster_identifier` (String)
- `final_db_snapshot_identifier` (String)
- `region` (String)
- `skip_final_snapshot` (String)
- `specify_rds_cluster` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--delete_rds_instance_action_value"></a>
### Nested Schema for `delete_rds_instance_action_value`

Read-Only:

- `final_rds_snapshot_id` (String)
- `rds_instance_id` (String)
- `region` (String)
- `skip_final_rds_snapshot` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--deregister_instances_action_value"></a>
### Nested Schema for `deregister_instances_action_value`

Read-Only:

- `instance_id` (String)
- `load_balancer_name` (String)
- `region` (String)
- `specify_instance` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--deregister_target_instances_action_value"></a>
### Nested Schema for `deregister_target_instances_action_value`

Read-Only:

- `region` (String)
- `tag_key` (String)
- `tag_value` (String)
- `target_group_arn` (String)


<a id="nestedatt--reboot_rds_instances_action_value"></a>
### Nested Schema for `reboot_rds_instances_action_value`

Read-Only:

- `rds_instance_id` (String)
- `region` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--reboot_workspaces_action_value"></a>
### Nested Schema for `reboot_workspaces_action_value`

Read-Only:

- `region` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--register_instances_action_value"></a>
### Nested Schema for `register_instances_action_value`

Read-Only:

- `instance_id` (String)
- `load_balancer_name` (String)
- `region` (String)
- `specify_instance` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--register_target_instances_action_value"></a>
### Nested Schema for `register_target_instances_action_value`

Read-Only:

- `region` (String)
- `tag_key` (String)
- `tag_value` (String)
- `target_group_arn` (String)


<a id="nestedatt--restore_from_cluster_snapshot_action_value"></a>
### Nested Schema for `restore_from_cluster_snapshot_action_value`

Read-Only:

- `allow_version_upgrade` (String)
- `availability_zone` (String)
- `cluster_identifier` (String)
- `cluster_parameter_group_name` (String)
- `cluster_subnet_group_name` (String)
- `delete_cluster_snapshot` (String)
- `port` (Number)
- `publicly_accessible` (String)
- `region` (String)
- `snapshot_identifier` (String)
- `vpc_security_group_ids` (List of String)


<a id="nestedatt--restore_rds_cluster_action_value"></a>
### Nested Schema for `restore_rds_cluster_action_value`

Read-Only:

- `auto_minor_version_upgrade` (String)
- `availability_zone` (String)
- `db_cluster_identifier` (String)
- `db_cluster_parameter_group_name` (String)
- `db_instance_class` (String)
- `db_instance_identifier` (String)
- `db_parameter_group_name` (String)
- `db_subnet_group_name` (String)
- `delete_db_cluster_snapshot` (String)
- `engine` (String)
- `engine_version` (String)
- `option_group_name` (String)
- `port` (Number)
- `publicly_accessible` (String)
- `region` (String)
- `snapshot_identifier` (String)
- `vpc_security_group_ids` (List of String)


<a id="nestedatt--restore_rds_instance_action_value"></a>
### Nested Schema for `restore_rds_instance_action_value`

Read-Only:

- `additional_tag_key` (String)
- `additional_tag_value` (String)
- `auto_minor_version_upgrade` (String)
- `availability_zone` (String)
- `db_engine` (String)
- `db_instance_class` (String)
- `db_name` (String)
- `delete_rds_snapshot` (String)
- `iops` (Number)
- `license_model` (String)
- `multi_az` (String)
- `option_group` (String)
- `parameter_group` (String)
- `port` (Number)
- `publicly_accessible` (String)
- `rds_instance_id` (String)
- `rds_snapshot_id` (String)
- `region` (String)
- `storage_type` (String)
- `subnet_group` (String)
- `trace_status` (String)
- `vpc` (String)
- `vpc_security_group_ids` (List of String)


<a id="nestedatt--revoke_security_group_ingress_action_value"></a>
### Nested Schema for `revoke_security_group_ingress_action_value`

Read-Only:

- `cidr_ip` (String)
- `ip_protocol` (String)
- `region` (String)
- `security_group_id` (String)
- `specify_security_group` (String)
- `tag_key` (String)
- `tag_value` (String)
- `to_port` (String)


<a id="nestedatt--schedule_rule_value"></a>
### Nested Schema for `schedule_rule_value`

Read-Only:

- `schedule` (String)
- `time_zone` (String)


<a id="nestedatt--send_command_action_value"></a>
### Nested Schema for `send_command_action_value`

Read-Only:

- `command` (String)
- `comment` (String)
- `document_name` (String)
- `execution_timeout_seconds` (Number)
- `instance_id` (String)
- `output_s3_bucket_name` (String)
- `output_s3_key_prefix` (String)
- `region` (String)
- `specify_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `timeout_seconds` (Number)
- `trace_status` (String)


<a id="nestedatt--sqs_v2_rule_value"></a>
### Nested Schema for `sqs_v2_rule_value`

Read-Only:

- `queue` (String)
- `sqs_aws_account_id` (Number)
- `sqs_region` (String)


<a id="nestedatt--start_instances_action_value"></a>
### Nested Schema for `start_instances_action_value`

Read-Only:

- `instance_id` (String)
- `region` (String)
- `specify_instance` (String)
- `status_checks_enable` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--start_rds_clusters_action_value"></a>
### Nested Schema for `start_rds_clusters_action_value`

Read-Only:

- `db_cluster_identifier` (String)
- `region` (String)
- `specify_rds_cluster` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--start_rds_instances_action_value"></a>
### Nested Schema for `start_rds_instances_action_value`

Read-Only:

- `rds_instance_id` (String)
- `region` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--start_workspaces_action_value"></a>
### Nested Schema for `start_workspaces_action_value`

Read-Only:

- `region` (String)
- `tag_key` (String)
- `tag_value` (String)


<a id="nestedatt--stop_instances_action_value"></a>
### Nested Schema for `stop_instances_action_value`

Read-Only:

- `instance_id` (String)
- `region` (String)
- `specify_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--stop_rds_clusters_action_value"></a>
### Nested Schema for `stop_rds_clusters_action_value`

Read-Only:

- `db_cluster_identifier` (String)
- `region` (String)
- `specify_rds_cluster` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--stop_rds_instances_action_value"></a>
### Nested Schema for `stop_rds_instances_action_value`

Read-Only:

- `rds_instance_id` (String)
- `region` (String)
- `specify_rds_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--terminate_workspaces_action_value"></a>
### Nested Schema for `terminate_workspaces_action_value`

Read-Only:

- `region` (String)
- `specify_workspace` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)


<a id="nestedatt--update_record_set_action_value"></a>
### Nested Schema for `update_record_set_action_value`

Read-Only:

- `record_set_name` (String)
- `record_set_type` (String)
- `record_set_value` (String)
- `zone_name` (String)


<a id="nestedatt--windows_update_action_value"></a>
### Nested Schema for `windows_update_action_value`

Read-Only:

- `comment` (String)
- `document_name` (String)
- `instance_id` (String)
- `kb_article_ids` (String)
- `output_s3_bucket_name` (String)
- `output_s3_key_prefix` (String)
- `region` (String)
- `specify_instance` (String)
- `tag_key` (String)
- `tag_value` (String)
- `timeout_seconds` (Number)
- `update_level` (String)


<a id="nestedatt--windows_update_v2_action_value"></a>
### Nested Schema for `windows_update_v2_action_value`

Read-Only:

- `allow_reboot` (String)
- `instance_id` (String)
- `output_s3_bucket_name` (String)
- `output_s3_key_prefix` (String)
- `region` (String)
- `severity_levels` (List of String)
- `specify_instance` (String)
- `specify_severity` (String)
- `tag_key` (String)
- `tag_value` (String)
- `trace_status` (String)