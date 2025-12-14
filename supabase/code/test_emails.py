import os

# Caminhos
base_dir = os.path.dirname(os.path.abspath(__file__))
templates_dir = os.path.join(base_dir, 'volumes', 'auth', 'templates')
output_dir = os.path.join(base_dir, 'email_previews')
logo_path = os.path.join(base_dir, 'volumes', 'www', 'assets', 'img', 'logo.png')

# Garantir diretório de saída
os.makedirs(output_dir, exist_ok=True)

# Templates para testar
templates = ['recovery.html', 'magic_link.html', 'invite.html', 'email_change.html']

# Mock data
site_url_mock = "file:///" + os.path.join(base_dir, 'volumes', 'www').replace("\\", "/")
token_mock = "TEST_TOKEN_123"

print(f"Gerando previews em: {output_dir}")

for template_name in templates:
    input_path = os.path.join(templates_dir, template_name)
    output_path = os.path.join(output_dir, f"preview_{template_name}")
    
    if os.path.exists(input_path):
        with open(input_path, 'r', encoding='utf-8') as f:
            content = f.read()
            
        # Substituição simples de variáveis Go template
        preview_content = content.replace('{{ .SiteURL }}', site_url_mock)
        preview_content = preview_content.replace('{{ .Token }}', token_mock)
        
        with open(output_path, 'w', encoding='utf-8') as f:
            f.write(preview_content)
            
        print(f"Generated: {output_path}")
    else:
        print(f"Template not found: {input_path}")

print("\nVerificação concluída. Abra os arquivos gerados no navegador para validar o visual.")
