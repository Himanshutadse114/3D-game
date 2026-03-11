filepath = r'c:\Game\game.project'
with open(filepath, 'rb') as f:
    content = f.read()
    print(f"File size: {len(content)} bytes")
    print("Hex dump of 'main_collection' line:")
    lines = content.split(b'\n')
    for line in lines:
        if b'main_collection' in line:
            print(f"Raw: {line}")
            print(f"Hex: {line.hex()}")
            for b in line:
                print(f"'{chr(b)}' ({b})", end=' ')
            print()
