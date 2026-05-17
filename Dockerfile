# Stage 1: Install dependencies
FROM node:18-alpine AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app
# Only copy package.json since package-lock.json doesn't exist locally
COPY package.json ./
# Use install instead of ci
RUN npm install

# Stage 2: Build the app
FROM node:18-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .
# Fixed ENV syntax
ENV NEXT_TELEMETRY_DISABLED=1
RUN npm run build

# Stage 3: Production server
FROM node:18-alpine AS runner
WORKDIR /app
# Fixed ENV syntax
ENV NODE_ENV=production
ENV NEXT_TELEMETRY_DISABLED=1

# Create a non-root user
RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

# Copy necessary files from the builder stage
# COPY --from=builder /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE 3000
# Fixed ENV syntax
ENV PORT=3000
ENV HOSTNAME="0.0.0.0"

CMD ["node", "server.js"]