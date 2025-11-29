#!/usr/bin/env python3
import json
import os
import urllib.request
import re

# Create data directories
os.makedirs('data/cities', exist_ok=True)
os.makedirs('data/districts', exist_ok=True)
os.makedirs('data/villages', exist_ok=True)
os.makedirs('raw', exist_ok=True)

# Download wilayah.sql
print("Downloading wilayah.sql...")
url = "https://raw.githubusercontent.com/cahyadsn/wilayah/master/db/wilayah.sql"
urllib.request.urlretrieve(url, 'raw/wilayah.sql')

print("Processing wilayah.sql...")

# Data structures
states = []
cities = {}
districts = {}
villages = {}

# Read and process wilayah.sql
with open('raw/wilayah.sql', 'r', encoding='utf-8') as f:
    content = f.read()
    
    # Find all INSERT statements with VALUES
    # Pattern: ('code','name')
    pattern = r"\('([^']+)','([^']+)'\)"
    matches = re.findall(pattern, content)
    
    for code, name in matches:
        code = code.strip()
        name = name.strip()
        
        # Skip if not valid code format
        if not code.replace('.', '').isdigit():
            continue
        
        # Count dots to determine level
        dot_count = code.count('.')
        
        if dot_count == 0:
            # State/Province
            states.append({"code": code, "value": name})
            
        elif dot_count == 1:
            # City/Regency
            state_code = code.split('.')[0]
            if state_code not in cities:
                cities[state_code] = []
            cities[state_code].append({"code": code, "value": name})
            
        elif dot_count == 2:
            # District/Subdistrict
            parts = code.split('.')
            city_code = f"{parts[0]}.{parts[1]}"
            if city_code not in districts:
                districts[city_code] = []
            districts[city_code].append({"code": code, "value": name})
            
        elif dot_count == 3:
            # Village/Kelurahan
            parts = code.split('.')
            district_code = f"{parts[0]}.{parts[1]}.{parts[2]}"
            if district_code not in villages:
                villages[district_code] = []
            villages[district_code].append({"code": code, "value": name})

# Write states.json
print("Writing states.json...")
with open('data/states.json', 'w', encoding='utf-8') as f:
    json.dump(states, f, ensure_ascii=False, indent=2)

# Write city files
print(f"Writing {len(cities)} city files...")
for state_code, city_list in cities.items():
    with open(f'data/cities/{state_code}.json', 'w', encoding='utf-8') as f:
        json.dump(city_list, f, ensure_ascii=False, indent=2)

# Write district files
print(f"Writing {len(districts)} district files...")
for city_code, district_list in districts.items():
    with open(f'data/districts/{city_code}.json', 'w', encoding='utf-8') as f:
        json.dump(district_list, f, ensure_ascii=False, indent=2)

# Write village files
print(f"Writing {len(villages)} village files...")
for district_code, village_list in villages.items():
    with open(f'data/villages/{district_code}.json', 'w', encoding='utf-8') as f:
        json.dump(village_list, f, ensure_ascii=False, indent=2)

print("\nData extraction complete!")
print(f"States: {len(states)} entries")
print(f"Cities: {len(cities)} files")
print(f"Districts: {len(districts)} files")
print(f"Villages: {len(villages)} files")
