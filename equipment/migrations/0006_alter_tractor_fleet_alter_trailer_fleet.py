# Generated by Django 4.1.7 on 2023-03-13 14:09

from django.db import migrations, models
import django.db.models.deletion


class Migration(migrations.Migration):

    dependencies = [
        ("dispatch", "0005_alter_delaycode_code_alter_fleetcode_code_and_more"),
        ("equipment", "0005_trailer_trailer_unique_trailer_code_organization"),
    ]

    operations = [
        migrations.AlterField(
            model_name="tractor",
            name="fleet",
            field=models.ForeignKey(
                blank=True,
                help_text="Fleet of the equipment.",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="tractor",
                related_query_name="tractor",
                to="dispatch.fleetcode",
                verbose_name="Fleet",
            ),
        ),
        migrations.AlterField(
            model_name="trailer",
            name="fleet",
            field=models.ForeignKey(
                blank=True,
                help_text="Fleet of the trailer.",
                null=True,
                on_delete=django.db.models.deletion.CASCADE,
                related_name="trailer",
                to="dispatch.fleetcode",
                verbose_name="Fleet",
            ),
        ),
    ]
