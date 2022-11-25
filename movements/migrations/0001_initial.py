# Generated by Django 4.1.3 on 2022-11-25 01:48

from django.db import migrations, models
import django.db.models.deletion
import django_extensions.db.fields
import utils.models


class Migration(migrations.Migration):

    initial = True

    dependencies = [
        ("order", "0053_remove_hazardousmaterial_organization_and_more"),
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("worker", "0023_alter_workercomment_comment_type"),
        ("equipment", "0008_alter_equipment_aux_power_unit_type_and_more"),
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
    ]