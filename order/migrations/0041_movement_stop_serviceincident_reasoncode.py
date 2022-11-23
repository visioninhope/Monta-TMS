# Generated by Django 4.1.3 on 2022-11-23 18:30

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import utils.models


class Migration(migrations.Migration):

    dependencies = [
        ("worker", "0022_alter_worker_worker_type_and_more"),
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("dispatch", "0007_alter_dispatchcontrol_distance_method_and_more"),
        ("equipment", "0008_alter_equipment_aux_power_unit_type_and_more"),
        ("location", "0008_remove_locationcomment_location_lo_locatio_c52573_idx"),
        ("order", "0040_remove_qualifiercode_organization_and_more"),
    ]

    operations = [
        migrations.CreateModel(
            name="Movement",
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
                    "ref_num",
                    models.CharField(
                        editable=False,
                        help_text="Movement Reference Number",
                        max_length=10,
                        unique=True,
                        verbose_name="Movement Reference Number",
                    ),
                ),
                (
                    "status",
                    utils.models.ChoiceField(
                        choices=[
                            ("N", "New"),
                            ("P", "In Progress"),
                            ("C", "Completed"),
                            ("B", "Billed"),
                            ("V", "Voided"),
                        ],
                        default="N",
                        help_text="Status of the Movement",
                        max_length=1,
                        verbose_name="Status",
                    ),
                ),
                (
                    "equipment",
                    models.ForeignKey(
                        blank=True,
                        help_text="Equipment of the Movement",
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="movements",
                        related_query_name="movement",
                        to="equipment.equipment",
                        verbose_name="Equipment",
                    ),
                ),
                (
                    "order",
                    models.ForeignKey(
                        help_text="Order of the Movement",
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="movements",
                        related_query_name="movement",
                        to="order.order",
                        verbose_name="Order",
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
                    "primary_worker",
                    models.ForeignKey(
                        blank=True,
                        help_text="Primary Worker of the Movement",
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="movements",
                        related_query_name="movement",
                        to="worker.worker",
                        verbose_name="Primary Worker",
                    ),
                ),
                (
                    "secondary_worker",
                    models.ForeignKey(
                        blank=True,
                        help_text="Secondary Worker of the Movement",
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="secondary_movements",
                        related_query_name="secondary_movement",
                        to="worker.worker",
                        verbose_name="Secondary Worker",
                    ),
                ),
            ],
            options={
                "verbose_name": "Movement",
                "verbose_name_plural": "Movements",
            },
        ),
        migrations.CreateModel(
            name="Stop",
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
                    "status",
                    utils.models.ChoiceField(
                        choices=[
                            ("N", "New"),
                            ("P", "In Progress"),
                            ("C", "Completed"),
                            ("B", "Billed"),
                            ("V", "Voided"),
                        ],
                        default="N",
                        help_text="The status of the stop.",
                        max_length=1,
                    ),
                ),
                (
                    "sequence",
                    models.PositiveIntegerField(
                        default=1,
                        help_text="The sequence of the stop in the movement.",
                        verbose_name="Sequence",
                    ),
                ),
                (
                    "pieces",
                    models.PositiveIntegerField(
                        blank=True,
                        default=0,
                        help_text="Pieces",
                        null=True,
                        verbose_name="Pieces",
                    ),
                ),
                (
                    "weight",
                    models.PositiveIntegerField(
                        blank=True,
                        default=0,
                        help_text="Weight",
                        null=True,
                        verbose_name="Weight",
                    ),
                ),
                (
                    "address_line",
                    models.CharField(
                        help_text="Stop Address",
                        max_length=255,
                        verbose_name="Stop Address",
                    ),
                ),
                (
                    "appointment_time",
                    models.DateTimeField(
                        help_text="The time the equipment is expected to arrive at the stop.",
                        verbose_name="Stop Appointment Time",
                    ),
                ),
                (
                    "arrival_time",
                    models.DateTimeField(
                        blank=True,
                        help_text="The time the equipment actually arrived at the stop.",
                        null=True,
                        verbose_name="Stop Arrival Time",
                    ),
                ),
                (
                    "departure_time",
                    models.DateTimeField(
                        blank=True,
                        help_text="The time the equipment actually departed from the stop.",
                        null=True,
                        verbose_name="Stop Departure Time",
                    ),
                ),
                (
                    "stop_type",
                    utils.models.ChoiceField(
                        choices=[
                            ("P", "Pickup"),
                            ("SP", "Split Pickup"),
                            ("SD", "Split Drop Off"),
                            ("D", "Delivery"),
                            ("DO", "Drop Off"),
                        ],
                        help_text="The type of stop.",
                        max_length=2,
                    ),
                ),
                (
                    "location",
                    models.ForeignKey(
                        help_text="The location of the stop.",
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="stops",
                        related_query_name="stop",
                        to="location.location",
                        verbose_name="Location",
                    ),
                ),
                (
                    "movement",
                    models.ForeignKey(
                        help_text="The movement that the stop belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="stops",
                        related_query_name="stop",
                        to="order.movement",
                        verbose_name="Movement",
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
                "verbose_name": "Stop",
                "verbose_name_plural": "Stops",
                "ordering": ["movement", "sequence"],
            },
        ),
        migrations.CreateModel(
            name="ServiceIncident",
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
                    "delay_reason",
                    models.CharField(
                        blank=True, max_length=100, verbose_name="Delay Reason"
                    ),
                ),
                (
                    "delay_time",
                    models.DurationField(
                        blank=True, null=True, verbose_name="Delay Time"
                    ),
                ),
                (
                    "delay_code",
                    models.ForeignKey(
                        blank=True,
                        null=True,
                        on_delete=django.db.models.deletion.PROTECT,
                        related_name="service_incidents",
                        related_query_name="service_incident",
                        to="dispatch.delaycode",
                        verbose_name="Delay Code",
                    ),
                ),
                (
                    "movement",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="service_incidents",
                        related_query_name="service_incident",
                        to="order.movement",
                        verbose_name="Movement",
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
                    "stop",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="service_incidents",
                        related_query_name="service_incident",
                        to="order.stop",
                        verbose_name="Stop",
                    ),
                ),
            ],
            options={
                "verbose_name": "Service Incident",
                "verbose_name_plural": "Service Incidents",
            },
        ),
        migrations.CreateModel(
            name="ReasonCode",
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
                    "code",
                    models.CharField(
                        help_text="Code of the Reason Code",
                        max_length=255,
                        unique=True,
                        verbose_name="Code",
                    ),
                ),
                (
                    "code_type",
                    utils.models.ChoiceField(
                        choices=[("VOIDED", "Voided"), ("CANCELLED", "Cancelled")],
                        help_text="Code Type of the Reason Code",
                        max_length=9,
                        verbose_name="Code Type",
                    ),
                ),
                (
                    "description",
                    models.CharField(
                        help_text="Description of the Reason Code",
                        max_length=100,
                        verbose_name="Description",
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
                "verbose_name": "Reason Code",
                "verbose_name_plural": "Reason Codes",
                "ordering": ["code"],
            },
        ),
    ]
