# Generated by Django 4.1.7 on 2023-03-11 09:23

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("equipment", "0001_initial"),
    ]

    operations = [
        migrations.AlterField(
            model_name="equipment",
            name="code",
            field=models.CharField(
                help_text="Code of the equipment.", max_length=50, verbose_name="Code"
            ),
        ),
        migrations.AlterField(
            model_name="equipmentmaintenanceplan",
            name="name",
            field=models.CharField(
                help_text="Name of the equipment maintenance plan.",
                max_length=50,
                verbose_name="Name",
            ),
        ),
        migrations.AlterField(
            model_name="equipmentmanufacturer",
            name="name",
            field=models.CharField(
                help_text="Name of the equipment manufacturer.",
                max_length=50,
                verbose_name="Name",
            ),
        ),
        migrations.AlterField(
            model_name="equipmenttype",
            name="name",
            field=models.CharField(
                help_text="Name of the equipment type",
                max_length=50,
                verbose_name="Name",
            ),
        ),
        migrations.AddConstraint(
            model_name="equipment",
            constraint=models.UniqueConstraint(
                fields=("code", "organization"),
                name="unique_equipment_code_organization",
            ),
        ),
        migrations.AddConstraint(
            model_name="equipmentmaintenanceplan",
            constraint=models.UniqueConstraint(
                fields=("name", "organization"),
                name="unique_equipment_maintenance_plan_name_organization",
            ),
        ),
        migrations.AddConstraint(
            model_name="equipmentmanufacturer",
            constraint=models.UniqueConstraint(
                fields=("name", "organization"),
                name="unique_equipment_manufacturer_organization",
            ),
        ),
        migrations.AddConstraint(
            model_name="equipmenttype",
            constraint=models.UniqueConstraint(
                fields=("name", "organization"),
                name="unique_equipment_type_name_organization",
            ),
        ),
    ]