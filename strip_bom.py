import os

extensions = {'.collection', '.gui', '.gui_script', '.script', '.render', '.render_script'}
root_dir = r'c:\Game'

for root, dirs, files in os.walk(root_dir):
    for file in files:
        if os.path.splitext(file)[1] in extensions:
            path = os.path.join(root, file)
            try:
                with open(path, 'rb') as f:
                    content = f.read()
                
                if content.startswith(b'\xef\xbb\xbf'):
                    print(f"Stripping BOM from {path}")
                    with open(path, 'wb') as f:
                        f.write(content[3:])
            except Exception as e:
                print(f"Error processing {path}: {e}")
