
import React from 'react';
import { withRorkErrorBoundary } from './rork-error-boundary';
import { App } from 'expo-router/build/qualified-entry';
import { renderRootComponent } from 'expo-router/build/renderRootComponent';

renderRootComponent(withRorkErrorBoundary(App));
