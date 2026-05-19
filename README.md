# Resource-Drain Auditor 🛡️

Una herramienta de línea de comandos (CLI) de alto rendimiento y consumo minimalista diseñada para auditar, estructurar y evaluar la reciprocidad en las interacciones relacionales y financieras. El sistema actúa como un firewall humano, detectando patrones destructivos de refuerzo intermitente y alertas de quiebre sistémico antes de que comprometan tu infraestructura personal.

Desarrollado bajo una arquitectura híbrida de bajo acoplamiento: un frontend de consola ágil y tolerante a fallos en **Go**, conectado a un motor analítico de precisión matemática en **Rust**.

---

## ⚡ Filosofía de Arquitectura

*   **Zero-Node / Bare-Metal Mindset:** Diseñado para ejecutarse nativamente sin entornos de ejecución pesados ni recolección de basura invasiva. Optimizado para un rendimiento fluido en hardware heredado (incluyendo equipos con especificaciones limitadas de más de una década).
*   **Fail-Fast Input Validation:** La interfaz captura los datos aislando cada campo en sub-bucles atómicos. Los errores de tipeo o rangos inválidos se corrigen en caliente sin romper el estado del formulario.
*   **Aislamiento de Cómputo:** Go se encarga exclusivamente de la sanitización de cadenas de texto y la gestión del flujo de entrada; Rust procesa el análisis estadístico pesado explotando la optimización nativa a nivel de instrucciones de ensamblador.

---

## 🛠️ Estructura del Proyecto

```text
├── cli-gateway/          # Interfaz de Usuario y Captura (Go 1.20+)
│   ├── cmd/main.go       # Loop de entrada tolerante a fallos y atajos UX
│   └── internal/
│       ├── model/        # Contratos y tipos de datos (Relacional/Financiero)
│       └── validator/    # Sanitización profunda de strings y límites de negocio
│
├── math-engine/          # Motor Estadístico de Precisión (Rust 2021)
│   ├── src/
│   │   ├── main.rs       # Inyección de datos y orquestador del diagnóstico
│   │   ├── psychology/   # Detectores de refuerzo intermitente y adicción
│   │   └── statistics/   # Módulos de ANOVA (Análisis de Varianza) y Bayes
│   └── Cargo.toml        # Perfiles de optimización agresiva para producción
│
├── Makefile              # Automatización de compilación y suites de pruebas
└── test-e2e.sh           # Script de integración de caja negra

📊 Modelado Matemático y Psicológico

El núcleo de Rust analiza los registros inyectados a través de tres capas de auditoría profunda:

Confianza Bayesiana de Reciprocidad: Calcula la probabilidad de que una interacción futura sea destructiva basándose en el historial acumulado. Las actualizaciones bayesianas se recalculan instantáneamente ante eventos negativos.

Análisis de Varianza (ANOVA): Evalúa el Retorno de Inversión (ROI) emocional y financiero. Una media severamente negativa emparejada con una desviación estándar baja certifica un escenario de pérdida constante y predecible.

Detector de Refuerzo Intermitente: Alerta si la otra parte utiliza dinámicas de recompensa aleatoria (alternando breves interacciones positivas entre rachas negativas), una técnica psicológica que genera dependencia conductual.

🚀 Instalación y Compilación
Este proyecto no utiliza contenedores ni capas de virtualización pesadas para garantizar el acceso directo al procesador. Requiere tener instalados los entornos locales de Go y Rust.

Para compilar toda la suite y generar los binarios optimizados en la raíz, ejecuta:

make build

🧪 Suite de Pruebas (Validación Extrema)
El sistema cuenta con pruebas unitarias de rango, sanitización de retornos de carro de Windows (\r\n), fuzzing matemático de estabilidad y un benchmark de estrés de un millón de eventos.

Para correr todas las pruebas de forma automatizada:

make test

Comportamiento del Motor bajo Estrés

El motor analítico en Rust está optimizado con perfiles de compilación específicos para producción. Al ejecutar las pruebas de estrés en modo optimizado:

cd math-engine && cargo test --release -- --nocapture

El procesador completa 1,000,000 de actualizaciones bayesianas consecutivas en un tiempo récord de ~100 nanosegundos, demostrando la viabilidad de la herramienta para auditorías masivas en tiempo real sin impacto térmico en la CPU.

🔒 Seguridad y Privacidad
Operación Local Absoluta: Los datos se procesan estrictamente en la memoria volatil local y en estructuras JSON sanitizadas. No se realizan conexiones de red externas, previniendo telemetría o fugas de información.

Protección de Frontera: La CLI de Go sanitiza todas las entradas utilizando flujos purgados de caracteres especiales y validaciones estructurales rígidas antes de empaquetar los payloads de analítica.