import uuid

import django.db.models.deletion
import django_extensions.db.fields
from django.db import migrations, models


class Migration(migrations.Migration):
    dependencies = [
        ("organization", "0003_alter_organization_phone_number"),
    ]

    operations = [
        migrations.CreateModel(
            name="TableChangeAlert",
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
                    ),
                ),
                (
                    "name",
                    models.CharField(
                        help_text="The name of the table change alert.",
                        max_length=50,
                        verbose_name="Name",
                    ),
                ),
                (
                    "table",
                    models.CharField(
                        choices=[
                            ("d", "d"),
                            ("j", "j"),
                            ("a", "a"),
                            ("n", "n"),
                            ("g", "g"),
                            ("o", "o"),
                            ("_", "_"),
                            ("m", "m"),
                            ("i", "i"),
                            ("g", "g"),
                            ("r", "r"),
                            ("a", "a"),
                            ("t", "t"),
                            ("i", "i"),
                            ("o", "o"),
                            ("n", "n"),
                            ("s", "s"),
                        ],
                        help_text="The table that the table change alert is for.",
                        max_length=50,
                        verbose_name="Table",
                    ),
                ),
                (
                    "organization",
                    models.ForeignKey(
                        help_text="The organization that the tax rate belongs to.",
                        on_delete=django.db.models.deletion.CASCADE,
                        related_name="table_change_alerts",
                        to="organization.organization",
                        verbose_name="Organization",
                    ),
                ),
            ],
            options={
                "verbose_name": "Table Change Alert",
                "verbose_name_plural": "Table Change Alerts",
            },
        ),
    ]
