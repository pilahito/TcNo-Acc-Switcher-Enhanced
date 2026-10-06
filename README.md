# TcNo Account Switcher Mejorado

Este repositorio es un fork mejorado del TcNo Account Switcher oficial, diseñado como un espacio de trabajo para construir una versión más robusta y con actualizaciones automáticas.

**Versión base utilizada:**
- https://github.com/TCNOco/TcNo-Acc-Switcher/releases/tag/2025-11-20_03

## Objetivo del Fork

El propósito de este fork es agregar las siguientes características al aplicativo original:

- ✅ **Actualizar la aplicación desde la propia aplicación sin descargar manualmente**
- ✅ Actualizar todas las cuentas configuradas en una sola acción
- ✅ Mejor validación de lanzamientos y seguridad de reversión
- ✅ Reporte de progreso más limpio y resumen consolidado
- ✅ Solución de problemas mejorada para Epic/Battle.net/Discord/Ubisoft

## Características Modeladas

- `CheckRelease` - Validación de versiones upstream
- `UpdateAll` - Lógica de actualización por lotes
- CLI demo para verificar versiones y procesar todas las cuentas
- Estructura orientada a lanzamientos, adecuada para un primer fork público

## Inicio Rápido

```bash
# Verificar nuevas versiones
go run . --check-release

# Actualizar todas las cuentas
go run . --update-all
```

## Estructura del Proyecto

```
cmd/                 Integración CLI futura
internal/
  accounts/          Gestión de cuentas
  updatecheck/       Verificación de actualizaciones
main.go              Punto de entrada de demo
README.md
CHANGELOG.md
```

## Plan de Lanzamiento

1. Alinear este fork con la última versión estable oficial
2. Agregar flujo de actualización interno en la aplicación (sin necesidad de descargar)
3. Agregar acciones de "Actualizar Todo" en la interfaz
4. Agregar verificación de versiones y seguridad de reversión
5. Publicar el primer lanzamiento etiquetado público

## Créditos

**Proyecto Original:**
- [TcNo Account Switcher](https://github.com/TCNOco/TcNo-Acc-Switcher) - Proyecto original por [TCNOco](https://github.com/TCNOco)

Este fork es un trabajo de mejora basado en el excelente trabajo del equipo de TcNo.

## Notas

Este repositorio está estructurado intencionalmente como una base de mejora robusta. No es el proyecto upstream; es una capa de mejora lista para fork construida alrededor del lanzamiento oficial.
