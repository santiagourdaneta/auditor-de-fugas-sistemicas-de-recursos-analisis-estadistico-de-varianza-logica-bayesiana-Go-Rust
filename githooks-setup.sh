#!/bin/bash
cat << 'EOF' > .git/hooks/pre-commit
#!/bin/bash
echo "🛡️  Validando calidad de código antes del commit..."
make lint || exit 1
make test || exit 1
echo "✅ Código perfecto. Procediendo al commit."
EOF
chmod +x .git/hooks/pre-commit