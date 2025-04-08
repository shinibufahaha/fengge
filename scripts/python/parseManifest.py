import xml.etree.ElementTree as ET
import json
import sys

class AndroidManifestParser:
    def __init__(self, file_path):
        self.file_path = file_path
        self.tree = None
        self.root = None

    def load_xml(self):
        try:
            self.tree = ET.parse(self.file_path)
            self.root = self.tree.getroot()
            print(f"Successfully loaded {self.file_path}")
        except ET.ParseError as e:
            print(f"Error parsing XML: {e}")

    def get_package_name(self):
        if self.root is None:
            print("XML not loaded. Please call load_xml() first.")
            return None

        return self.root.get('package')
    
    def get_component_data(self, component_name):
        if self.root is None:
            print("XML not loaded. Please call load_xml() first.")
            return None

        components = self.root.findall(f".//{component_name}")
        component_data = []

        for component in components:
            component_info = {
                "name": component.get("{http://schemas.android.com/apk/res/android}name"),
                "exported": component.get("{http://schemas.android.com/apk/res/android}exported"),
                "enabled": component.get("{http://schemas.android.com/apk/res/android}enabled"),
                "permission": component.get("{http://schemas.android.com/apk/res/android}permission"),
                'readPermission': component.get("{http://schemas.android.com/apk/res/android}readPermission"),
                'writePermission': component.get("{http://schemas.android.com/apk/res/android}writePermission"),
                'authorities': component.get("{http://schemas.android.com/apk/res/android}authorities"),
            }

            intent_filters = component.findall("intent-filter")
            component_info["intent_filters"] = []

            for intent_filter in intent_filters:
                filter_info = {
                    "actions": [action.get("{http://schemas.android.com/apk/res/android}name") for action in intent_filter.findall("action")],
                    "categories": [category.get("{http://schemas.android.com/apk/res/android}name") for category in intent_filter.findall("category")],
                    "data": [{
                        "scheme": data.get("{http://schemas.android.com/apk/res/android}scheme"),
                        "host": data.get("{http://schemas.android.com/apk/res/android}host"),
                        "path": data.get("{http://schemas.android.com/apk/res/android}path"),
                        "pathPrefix": data.get("{http://schemas.android.com/apk/res/android}pathPrefix"),
                    } for data in intent_filter.findall("data")]
                }
                component_info["intent_filters"].append(filter_info)

            component_data.append(component_info)

        return component_data

    def save_to_json(self, data, output_file):
        try:
            with open(output_file, 'w', encoding='utf-8') as f:
                json.dump(data, f, ensure_ascii=False, indent=4)
            print(f"Data successfully saved to {output_file}")
        except IOError as e:
            print(f"Error saving to JSON: {e}")

if __name__ == "__main__":
    file_path = sys.argv[1]
    output_file = sys.argv[2]

    parser = AndroidManifestParser(file_path)
    parser.load_xml()

    data = {
        "activities": parser.get_component_data("activity"),
        "services": parser.get_component_data("service"),
        "receivers": parser.get_component_data("receiver"),
        "providers": parser.get_component_data("provider"),
    }
    
    parser.save_to_json(data, output_file)