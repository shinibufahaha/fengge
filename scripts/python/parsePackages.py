import xml.etree.ElementTree as ET
import json

def parse_permissions(in_path, out_path):
    tree = ET.parse(in_path)
    root = tree.getroot()
    json_data = {}

    for item in root.findall(".//permissions/item"):
        name = item.get('name')
        protection = item.get('protection')
        if protection is None:
            protection = "0"
        json_data[name] = protection 
    with open(out_path, 'w') as json_file:
        json.dump(json_data, json_file, indent=4)
