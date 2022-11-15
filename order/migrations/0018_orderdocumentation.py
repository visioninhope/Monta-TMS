# Generated by Django 4.1.3 on 2022-11-15 06:20

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models

import order.models


class Migration(migrations.Migration):

    dependencies = [
        ("billing", "0015_alter_accessorialcharge_options_and_more"),
        ("organization", "0007_alter_depot_description_alter_depot_name_and_more"),
        ("order", "0017_stop_serviceincident"),
    ]

    operations = [
        migrations.CreateModel(
            name="OrderDocumentation",
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
                    "document",
                    models.FileField(
                        blank=True,
                        null=True,
                        upload_to=order.models.order_documentation_upload_to,
                        verbose_name="Document",
                    ),
                ),
                (
                    "document_class",
                    models.ForeignKey(
                        help_text="Document Class",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="order_documentation",
                        to="billing.documentclassification",
                        verbose_name="Document Class",
                    ),
                ),
                (
                    "order",
                    models.ForeignKey(
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="order_documentation",
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
            ],
            options={
                "verbose_name": "Order Documentation",
                "verbose_name_plural": "Order Documentation",
            },
        ),
    ]
