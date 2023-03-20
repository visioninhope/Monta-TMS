# Generated by Django 4.1.7 on 2023-03-19 04:56

from django.conf import settings
from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import uuid


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        migrations.swappable_dependency(settings.AUTH_USER_MODEL),
        ("organization", "0015_alter_tablechangealert_table"),
    ]

    operations = [
        migrations.CreateModel(
            name="CustomReport",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                        verbose_name="ID",
                    ),
                ),
                (
                    "name",
                    models.CharField(max_length=255, unique=True, verbose_name="Name"),
                ),
                (
                    "table",
                    models.CharField(
                        choices=[
                            ("additional_charge", "additional_charge"),
                            ("auditlog_logentry", "auditlog_logentry"),
                            ("auth_group_permissions", "auth_group_permissions"),
                            ("auth_token", "auth_token"),
                            ("billing_control", "billing_control"),
                            ("billing_exception", "billing_exception"),
                            ("billing_history", "billing_history"),
                            ("billing_queue", "billing_queue"),
                            ("billing_transfer_log", "billing_transfer_log"),
                            ("charge_type", "charge_type"),
                            ("comment_type", "comment_type"),
                            ("commodity", "commodity"),
                            ("custom_report", "custom_report"),
                            ("customer", "customer"),
                            ("customer_billing_profile", "customer_billing_profile"),
                            ("customer_contact", "customer_contact"),
                            ("customer_email_profile", "customer_email_profile"),
                            ("customer_fuel_profile", "customer_fuel_profile"),
                            (
                                "customer_fuel_profile_detail",
                                "customer_fuel_profile_detail",
                            ),
                            ("customer_fuel_table", "customer_fuel_table"),
                            ("customer_rule_profile", "customer_rule_profile"),
                            (
                                "customer_rule_profile_document_class",
                                "customer_rule_profile_document_class",
                            ),
                            ("delay_code", "delay_code"),
                            ("department", "department"),
                            ("depot", "depot"),
                            ("depot_detail", "depot_detail"),
                            ("dispatch_control", "dispatch_control"),
                            ("division_code", "division_code"),
                            (
                                "django_celery_beat_clockedschedule",
                                "django_celery_beat_clockedschedule",
                            ),
                            (
                                "django_celery_beat_intervalschedule",
                                "django_celery_beat_intervalschedule",
                            ),
                            (
                                "django_celery_beat_periodictasks",
                                "django_celery_beat_periodictasks",
                            ),
                            (
                                "django_celery_results_chordcounter",
                                "django_celery_results_chordcounter",
                            ),
                            (
                                "django_celery_results_taskresult",
                                "django_celery_results_taskresult",
                            ),
                            ("django_migrations", "django_migrations"),
                            ("document_classification", "document_classification"),
                            ("email_control", "email_control"),
                            ("email_log", "email_log"),
                            ("email_profile", "email_profile"),
                            (
                                "equipment_maintenance_plan",
                                "equipment_maintenance_plan",
                            ),
                            (
                                "equipment_maintenance_plan_equipment_types",
                                "equipment_maintenance_plan_equipment_types",
                            ),
                            ("equipment_manufacturer", "equipment_manufacturer"),
                            ("equipment_type", "equipment_type"),
                            ("equipment_type_detail", "equipment_type_detail"),
                            ("fleet_code", "fleet_code"),
                            ("fuel_vendor", "fuel_vendor"),
                            ("fuel_vendor_fuel_detail", "fuel_vendor_fuel_detail"),
                            ("general_ledger_account", "general_ledger_account"),
                            ("google_api", "google_api"),
                            ("hazardous_material", "hazardous_material"),
                            ("integration", "integration"),
                            ("integration_vendor", "integration_vendor"),
                            ("invoice_control", "invoice_control"),
                            ("job_title", "job_title"),
                            ("location", "location"),
                            ("location_category", "location_category"),
                            ("location_comment", "location_comment"),
                            ("location_contact", "location_contact"),
                            ("movement", "movement"),
                            ("order", "order"),
                            ("order_comment", "order_comment"),
                            ("order_control", "order_control"),
                            ("order_documentation", "order_documentation"),
                            ("order_type", "order_type"),
                            ("organization", "organization"),
                            ("other_charge", "other_charge"),
                            ("qualifier_code", "qualifier_code"),
                            ("rate", "rate"),
                            ("rate_billing_table", "rate_billing_table"),
                            ("rate_table", "rate_table"),
                            ("reason_code", "reason_code"),
                            ("report_column", "report_column"),
                            ("revenue_code", "revenue_code"),
                            ("route", "route"),
                            ("route_control", "route_control"),
                            ("scheduled_report", "scheduled_report"),
                            (
                                "scheduled_report_day_of_week",
                                "scheduled_report_day_of_week",
                            ),
                            ("service_incident", "service_incident"),
                            ("silk_profile_queries", "silk_profile_queries"),
                            ("silk_response", "silk_response"),
                            ("stop", "stop"),
                            ("stop_comment", "stop_comment"),
                            ("table_change_alert", "table_change_alert"),
                            ("tax_rate", "tax_rate"),
                            ("tractor", "tractor"),
                            ("trailer", "trailer"),
                            ("user", "user"),
                            ("user_groups", "user_groups"),
                            ("user_profile", "user_profile"),
                            ("user_user_permissions", "user_user_permissions"),
                            ("weekday", "weekday"),
                            ("worker", "worker"),
                            ("worker_comment", "worker_comment"),
                            ("worker_contact", "worker_contact"),
                            ("worker_profile", "worker_profile"),
                            ("worker_time_away", "worker_time_away"),
                        ],
                        help_text="The table that the table change alert is for.",
                        max_length=255,
                        verbose_name="Table",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Custom Report",
                "verbose_name_plural": "Custom Reports",
                "db_table": "custom_report",
                "ordering": ("name",),
            },
        ),
        migrations.CreateModel(
            name="Weekday",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        choices=[
                            (0, "Monday"),
                            (1, "Tuesday"),
                            (2, "Wednesday"),
                            (3, "Thursday"),
                            (4, "Friday"),
                            (5, "Saturday"),
                            (6, "Sunday"),
                        ],
                        max_length=15,
                        unique=True,
                        verbose_name="Name",
                    ),
                ),
            ],
            options={
                "verbose_name": "Weekday",
                "verbose_name_plural": "Weekdays",
                "db_table": "weekday",
                "ordering": ("name",),
            },
        ),
        migrations.CreateModel(
            name="ScheduledReport",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        help_text="Unique ID for the user.",
                        primary_key=True,
                        serialize=False,
                        unique=True,
                    ),
                ),
                (
                    "is_active",
                    models.BooleanField(
                        default=True,
                        help_text="Whether the scheduled report is active.",
                        verbose_name="Active",
                    ),
                ),
                (
                    "schedule_type",
                    models.CharField(
                        choices=[
                            ("DAILY", "Daily"),
                            ("WEEKLY", "Weekly"),
                            ("MONTHLY", "Monthly"),
                        ],
                        default="DAILY",
                        max_length=255,
                        verbose_name="Schedule Type",
                    ),
                ),
                (
                    "time",
                    models.TimeField(
                        help_text="The time of day to send the report.",
                        verbose_name="Time",
                    ),
                ),
                (
                    "day_of_month",
                    models.PositiveIntegerField(
                        blank=True,
                        help_text="The day of the month to send the report.",
                        null=True,
                        verbose_name="Day of Month",
                    ),
                ),
                (
                    "day_of_week",
                    models.ManyToManyField(
                        blank=True,
                        help_text="The day of the week to send the report.",
                        to="reports.weekday",
                        verbose_name="Day of Week",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "report",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="scheduled_reports",
                        to="reports.customreport",
                        verbose_name="Report",
                    ),
                ),
                (
                    "user",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="scheduled_reports",
                        to=settings.AUTH_USER_MODEL,
                        verbose_name="User",
                    ),
                ),
            ],
            options={
                "verbose_name": "Scheduled Report",
                "verbose_name_plural": "Scheduled Reports",
                "db_table": "scheduled_report",
                "ordering": ("report",),
            },
        ),
        migrations.CreateModel(
            name="ReportColumn",
            fields=[
                (
                    "created",
                    django_extensions.db.fields.CreationDateTimeField(
                        auto_now_add=True, verbose_name="created"
                    ),
                ),
                (
                    "modified",
                    django_extensions.db.fields.ModificationDateTimeField(
                        auto_now=True, verbose_name="modified"
                    ),
                ),
                (
                    "id",
                    models.UUIDField(
                        default=uuid.uuid4,
                        editable=False,
                        primary_key=True,
                        serialize=False,
                        unique=True,
                        verbose_name="ID",
                    ),
                ),
                (
                    "column_name",
                    models.CharField(
                        help_text="The name of the column to be displayed in the report.",
                        max_length=255,
                        verbose_name="Column Name",
                    ),
                ),
                (
                    "column_order",
                    models.PositiveIntegerField(
                        help_text="The order of the column to be displayed in the report.",
                        verbose_name="Column Order",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="Organization",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="%(class)ss",
                        related_query_name="%(class)s",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
                (
                    "report",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="columns",
                        to="reports.customreport",
                        verbose_name="Report",
                    ),
                ),
            ],
            options={
                "verbose_name": "Report Column",
                "verbose_name_plural": "Report Columns",
                "db_table": "report_column",
                "ordering": ("column_order",),
            },
        ),
        migrations.AddConstraint(
            model_name="scheduledreport",
            constraint=models.UniqueConstraint(
                fields=("report", "organization"),
                name="unique_scheduled_report_report_organization",
            ),
        ),
        migrations.AddConstraint(
            model_name="reportcolumn",
            constraint=models.UniqueConstraint(
                fields=("report", "column_name", "column_order"),
                name="unique_report_column_name_order",
            ),
        ),
        migrations.AddConstraint(
            model_name="customreport",
            constraint=models.UniqueConstraint(
                fields=("name", "organization"), name="unique_report_name_organization"
            ),
        ),
    ]
