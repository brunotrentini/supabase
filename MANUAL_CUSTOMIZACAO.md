# Manual de Customização de Identidade (White-Label)

Este documento orienta como personalizar a identidade visual, logos e templates de e-mail da instância do Supabase para adequá-la a um novo SaaS ou cliente específico (White-label).

---

## 1. Identidade Visual (Logos)

As imagens utilizadas tanto nas páginas web (Login, Recuperação) quanto nos templates de e-mail estão centralizadas em um único diretório acessível publicamente.

**Diretório:** `supabase/code/volumes/www/assets/img/`

Para alterar a marca, você tem duas opções:

### Opção A: Substituição Direta (Recomendado)
Substitua os arquivos existentes mantendo EXATAMENTE os mesmos nomes. Isso evita a necessidade de alterar códigos HTML.

1.  **`logo.png`**: Logo colorida/principal (fundo claro).
    *   *Uso*: E-mails, Tela de Login, Recuperação de Senha.
    *   *Dimensão Sugerida*: Altura de ~50px a 80px (PNG transparente).
2.  **`logo-white.png`**: Logo branca/negativa (fundo escuro).
    *   *Uso*: Telas com fundo escuro (se houver).

### Opção B: Novos Arquivos
Se adicionar arquivos com nomes diferentes (ex: `logo-cliente-x.png`), você precisará editar todos os arquivos HTML listados na seção 2 e 3 para apontar para o novo caminho.

---

## 2. Templates de E-mail (HTML)

Os textos e o layout dos e-mails transacionais podem ser editados livremente, desde que as variáveis de sistema sejam mantidas.

**Diretório:** `supabase/code/volumes/auth/templates/`

### Arquivos Principais:
*   `recovery.html`: E-mail de "Esqueci minha senha".
*   `magic_link.html`: E-mail de "Link de acesso sem senha" (Magic Link).
*   `invite.html`: E-mail de convite para novos membros.
*   `email_change.html`: E-mail de confirmação de troca de e-mail.

### ⚠️ Regras Importantes de Edição:
Ao editar o HTML, **NÃO REMOVA** as seguintes variáveis (elas são preenchidas automaticamente pelo sistema):
*   `{{ .SiteURL }}`: A URL base do seu sistema (ex: https://auth.seusistema.com).
*   `{{ .Token }}`: O token de segurança único do usuário.
*   `{{ .ConfirmationURL }}`: (Se usado) O link completo de confirmação.

**Exemplo de Link Seguro:**
```html
<a href="{{ .SiteURL }}/recovery-confirm?token={{ .Token }}&type=recovery">
    Redefinir Senha
</a>
```

---

## 3. Configuração de Assuntos e Remetente (.env)

O **Assunto do E-mail (Subject)** e o **Remetente** não ficam no HTML. Eles são configurados nas variáveis de ambiente.

**Arquivo:** `supabase/code/.env`

Edite as seguintes linhas para personalizar:

```env
## Configurações de SMTP (Quem envia)
SMTP_ADMIN_EMAIL=suporte@novocliente.com
SMTP_SENDER_NAME="Nome do Novo SaaS"

## Assuntos dos E-mails (Subjects)
GOTRUE_MAILER_SUBJECTS_RECOVERY="Redefina sua senha - Novo SaaS"
GOTRUE_MAILER_SUBJECTS_MAGIC_LINK="Seu link de acesso - Novo SaaS"
GOTRUE_MAILER_SUBJECTS_INVITE="Convite para acessar o Novo SaaS"
GOTRUE_MAILER_SUBJECTS_CONFIRMATION="Confirme seu cadastro - Novo SaaS"
GOTRUE_MAILER_SUBJECTS_EMAIL_CHANGE="Troca de e-mail - Novo SaaS"
```

---

## 4. Páginas Web do Auth

Além dos e-mails, existem as páginas HTML que o usuário vê ao clicar no link (ex: página para digitar a nova senha).

**Diretório:** `supabase/code/volumes/www/`

*   `recovery.html`: Página onde o usuário digita a nova senha.
*   `welcome.html`: Página de sucesso/boas-vindas.
*   `*.html`: Outras páginas de suporte.

**O que editar:**
*   Textos de boas-vindas.
*   Cores dos botões (busque por classes Tailwind como `bg-blue-600` e troque pela cor da marca, ex: `bg-green-600`).

---

## 5. Aplicando as Alterações

Após modificar qualquer arquivo (especialmente o `.env`), é necessário reiniciar o serviço de autenticação para que as mudanças tenham efeito.

**Comando (no terminal):**
```bash
docker compose restart auth
```
*(Se alterou apenas arquivos HTML em `volumes/`, a atualização costuma ser imediata, mas reiniciar garante a limpeza de cache).*

---

## 6. Como Testar

Para visualizar como os e-mails ficarão antes de enviar para um usuário real:

1.  Abra o terminal na pasta `supabase/code`.
2.  Execute o script de teste:
    ```bash
    python test_emails.py
    ```
3.  Abra a pasta `email_previews` criada e visualize os arquivos `.html` no seu navegador.
